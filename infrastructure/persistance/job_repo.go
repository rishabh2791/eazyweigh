package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"errors"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobRepo struct {
	DB          *gorm.DB
	WarehouseDB *gorm.DB
	Logger      hclog.Logger
}

var _ repository.JobRepository = &JobRepo{}

func NewJobRepo(db *gorm.DB, warehouseDB *gorm.DB, logger hclog.Logger) *JobRepo {
	return &JobRepo{
		DB:          db,
		WarehouseDB: warehouseDB,
		Logger:      logger,
	}
}

func getJobCode(jobCode string) string {
	code := ""
	for i := 0; i < 15-len(jobCode); i++ {
		code += "0"
	}
	return code + jobCode
}

type RemoteMaterial struct {
	StockCode   string
	Description string
}

type RemoteBOMItem struct {
	StockCode string
	Quantity  float32
}

func (jobRepo *JobRepo) getMaterialFromRemote(stockCode string) (*RemoteMaterial, error) {
	remoteMaterial := RemoteMaterial{}
	stockCodeQuery := "SELECT StockCode, Description FROM dbo.InvMaster WHERE StockCode = '" + stockCode + "'"
	rows, getErr := jobRepo.WarehouseDB.Raw(stockCodeQuery).Rows()
	defer rows.Close()
	if getErr != nil {
		return nil, getErr
	}
	for rows.Next() {
		scanErr := rows.Scan(&remoteMaterial.StockCode, &remoteMaterial.Description)
		if scanErr != nil {
			return nil, scanErr
		}
	}
	return &remoteMaterial, nil
}

func (jobRepo *JobRepo) getBOMFromRemote(stockCode string) ([]RemoteBOMItem, error) {
	remoteBOMItems := []RemoteBOMItem{}
	bomQuery := "SELECT Component, QtyPer FROM dbo.BomStructure WHERE ParentPart = '" + stockCode + "'"
	rows, getErr := jobRepo.WarehouseDB.Raw(bomQuery).Rows()
	defer rows.Close()
	if getErr != nil {
		return nil, getErr
	}
	for rows.Next() {
		remotBomItem := RemoteBOMItem{}
		scanErr := rows.Scan(&remotBomItem.StockCode, &remotBomItem.Quantity)
		if scanErr != nil {
			return nil, scanErr
		}
		remoteBOMItems = append(remoteBOMItems, remotBomItem)
	}
	return remoteBOMItems, nil
}

func (jobRepo *JobRepo) CreateMaterial(remoteMaterial *RemoteMaterial, job *entity.Job, ty string) (*entity.Material, error) {
	stock := entity.Material{}
	stock.Type = ty
	stock.FactoryID = job.FactoryID
	stock.Code = remoteMaterial.StockCode
	stock.Description = remoteMaterial.Description
	stock.UnitOfMeasureID = job.UnitOfMeasureID
	stock.CreatedByUsername = job.CreatedByUsername
	stock.UpdatedByUsername = job.UpdatedByUsername

	// Create Material
	stockCreationErr := jobRepo.DB.Create(&stock).Error
	if stockCreationErr != nil {
		return nil, stockCreationErr
	}
	return &stock, nil
}

func (jobRepo *JobRepo) CreateBOM(job *entity.Job, existingStockCode *entity.Material, stockCode string, revision int) (*entity.BOM, error) {
	remoteBOMItems, remoteErr := jobRepo.getBOMFromRemote(stockCode)
	if remoteErr != nil {
		return nil, remoteErr
	}
	bom := entity.BOM{}
	bom.ID = uuid.New().String()
	bom.FactoryID = job.FactoryID
	bom.CreatedByUsername = job.CreatedByUsername
	bom.UpdatedByUsername = job.UpdatedByUsername
	bom.MaterialID = existingStockCode.ID
	bom.UnitSize = 1
	bom.Revision = revision
	bom.UnitOfMeasureID = job.UnitOfMeasureID
	bomItems := []entity.BOMItem{}
	for _, remoteBOMItem := range remoteBOMItems {
		existingComponent := entity.Material{}
		getComponentErr := jobRepo.DB.Where("factory_id = ? AND code = ?", job.FactoryID, remoteBOMItem.StockCode).Take(&existingComponent).Error
		if getComponentErr != nil {
			//Not Created
			remoteComponent, remoteErr := jobRepo.getMaterialFromRemote(remoteBOMItem.StockCode)
			if remoteErr != nil {
				return nil, remoteErr
			}
			//Create Material
			component, getErr := jobRepo.CreateMaterial(remoteComponent, job, "Raw Material")
			if getErr != nil {
				return nil, getErr
			}
			existingComponent = *component
		}
		bomItem := entity.BOMItem{}
		bomItem.MaterialID = existingComponent.ID
		bomItem.Quantity = remoteBOMItem.Quantity
		bomItem.BOMID = bom.ID
		bomItem.UpperTolerance = 1
		bomItem.LowerTolerance = 1
		bomItem.UnitOfMeasureID = job.UnitOfMeasureID
		bomItem.CreatedByUsername = job.CreatedByUsername
		bomItem.UpdatedByUsername = job.UpdatedByUsername
		bomItems = append(bomItems, bomItem)
	}
	bom.BOMItems = bomItems
	creationErr := jobRepo.DB.Create(&bom).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return &bom, nil
}

func checkInRemoteBOMItems(stockCode string, bomItems []RemoteBOMItem) bool {
	for _, remoteBOMItem := range bomItems {
		if stockCode == remoteBOMItem.StockCode {
			return true
		}
	}
	return false
}

func (jobRepo *JobRepo) checkBOMRevision(stockCode string, bom *entity.BOM) (bool, error) {
	revision := true
	remoteBOMItems, remoteErr := jobRepo.getBOMFromRemote(stockCode)
	if remoteErr != nil {
		return true, remoteErr
	}
	for _, bomItem := range bom.BOMItems {
		check := checkInRemoteBOMItems(bomItem.Material.Code, remoteBOMItems)
		revision = revision && check
	}
	return revision, nil
}

func (jobRepo *JobRepo) Create(job *entity.Job) (*entity.Job, error) {
	var stockCode string
	var quantity float32

	// Get Material For which Job is Created.
	query := "SELECT StockCode, TrnQty FROM dbo.InvMovements WHERE Job = '" + getJobCode(job.JobCode) + "'"
	rows, getErr := jobRepo.WarehouseDB.Raw(query).Rows()

	if getErr != nil {
		return nil, getErr
	}
	defer rows.Close()

	for rows.Next() {
		var currentStockCode string
		var currentQuantity float32
		scanErr := rows.Scan(&currentStockCode, &currentQuantity)
		if scanErr != nil {
			return nil, scanErr
		}
		if currentStockCode[0] == 55 {
			stockCode = currentStockCode
			quantity = job.Quantity
		}
	}

	if stockCode != "" || len(stockCode) != 0 {
		// Check if material is created
		existingStockCode := entity.Material{}
		getMaterialErr := jobRepo.DB.Where("code = ? AND factory_id = ?", stockCode, job.FactoryID).Take(&existingStockCode).Error
		if getMaterialErr != nil {
			//Not Created
			remoteMaterial, remoteErr := jobRepo.getMaterialFromRemote(stockCode)
			if remoteErr != nil {
				return nil, remoteErr
			}

			//Create Material
			material, getErr := jobRepo.CreateMaterial(remoteMaterial, job, "Bulk")
			if getErr != nil {
				return nil, getErr
			}
			existingStockCode = *material
		}

		//Check if BOM Exists, If exists BOM Items are already created. There may be revisions
		existingBOM := entity.BOM{}
		getBomErr := jobRepo.DB.
			Preload("Material.").
			Preload("Material.UnitOfMeasure").
			Preload("Material.UnitOfMeasure.Factory").
			Preload("Material.UnitOfMeasure.Factory.Address").
			Preload("Material.UnitOfMeasure.Factory.CreatedBy").
			Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
			Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
			Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
			Preload("Material.UnitOfMeasure.CreatedBy").
			Preload("Material.UnitOfMeasure.UpdatedBy").
			Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
			Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
			Preload("Material.CreatedBy").
			Preload("Material.CreatedBy.UserRole").
			Preload("Material.UpdatedBy").
			Preload("Material.UpdatedBy.UserRole").
			Preload("BOMItems.Material.").
			Preload("BOMItems.Material.UnitOfMeasure").
			Preload("BOMItems.Material.UnitOfMeasure.Factory").
			Preload("BOMItems.Material.UnitOfMeasure.Factory.Address").
			Preload("BOMItems.Material.UnitOfMeasure.Factory.CreatedBy").
			Preload("BOMItems.Material.UnitOfMeasure.Factory.UpdatedBy").
			Preload("BOMItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
			Preload("BOMItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
			Preload("BOMItems.Material.UnitOfMeasure.CreatedBy").
			Preload("BOMItems.Material.UnitOfMeasure.UpdatedBy").
			Preload("BOMItems.Material.UnitOfMeasure.CreatedBy.UserRole").
			Preload("BOMItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
			Preload("BOMItems.Material.CreatedBy").
			Preload("BOMItems.Material.CreatedBy.UserRole").
			Preload("BOMItems.Material.UpdatedBy").
			Preload("BOMItems.Material.UpdatedBy.UserRole").
			Preload("BOMItems.UnitOfMeasure").
			Preload("BOMItems.UnitOfMeasure.Factory").
			Preload("BOMItems.UnitOfMeasure.Factory.Address").
			Preload("BOMItems.UnitOfMeasure.Factory.CreatedBy").
			Preload("BOMItems.UnitOfMeasure.Factory.UpdatedBy").
			Preload("BOMItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
			Preload("BOMItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
			Preload("BOMItems.UnitOfMeasure.CreatedBy").
			Preload("BOMItems.UnitOfMeasure.UpdatedBy").
			Preload("BOMItems.UnitOfMeasure.CreatedBy.UserRole").
			Preload("BOMItems.UnitOfMeasure.UpdatedBy.UserRole").
			Preload("BOMItems.CreatedBy").
			Preload("BOMItems.UpdatedBy").
			Preload("BOMItems.CreatedBy.UserRole").
			Preload("BOMItems.UpdatedBy.UserRole").
			Preload("UnitOfMeasure.Factory").
			Preload("UnitOfMeasure.Factory.Address").
			Preload("UnitOfMeasure.Factory.CreatedBy").
			Preload("UnitOfMeasure.Factory.UpdatedBy").
			Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
			Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
			Preload("UnitOfMeasure.CreatedBy").
			Preload("UnitOfMeasure.UpdatedBy").
			Preload("UnitOfMeasure.CreatedBy.UserRole").
			Preload("UnitOfMeasure.UpdatedBy.UserRole").
			Preload("CreatedBy.UserRole").
			Preload("UpdatedBy.UserRole").
			Preload(clause.Associations).Where("factory_id = ? AND material_id = ?", job.FactoryID, existingStockCode.ID).Take(&existingBOM).Error
		if getBomErr != nil {
			//Not Created
			bom, creationErr := jobRepo.CreateBOM(job, &existingStockCode, stockCode, 1)
			if creationErr != nil {
				return nil, creationErr
			}
			existingBOM = *bom
		} else {
			//BOM Already Created so check for revision
			bomRevision, err := jobRepo.checkBOMRevision(stockCode, &existingBOM)
			if err != nil {
				return nil, err
			}
			if !bomRevision {
				bom, creationErr := jobRepo.CreateBOM(job, &existingStockCode, stockCode, 2)
				if creationErr != nil {
					return nil, creationErr
				}
				existingBOM = *bom
			}
		}

		job.MaterialID = existingStockCode.ID
		job.ID = uuid.New().String()
		jobItems := []entity.JobItem{}

		for _, bomItem := range existingBOM.BOMItems {
			jobItem := entity.JobItem{}
			jobItem.JobID = job.ID
			jobItem.CreatedByUsername = job.CreatedByUsername
			jobItem.UpdatedByUsername = job.UpdatedByUsername
			jobItem.MaterialID = bomItem.MaterialID
			jobItem.RequiredWeight = bomItem.Quantity * quantity
			jobItem.LowerBound = bomItem.Quantity * quantity * (1.0 - bomItem.LowerTolerance/100)
			jobItem.UpperBound = bomItem.Quantity * quantity * (1.0 + bomItem.UpperTolerance/100)
			jobItem.UnitOfMeasureID = bomItem.UnitOfMeasureID
			jobItems = append(jobItems, jobItem)
		}

		job.JobItems = jobItems

		creationErr := jobRepo.DB.Create(&job).Error
		if creationErr != nil {
			return nil, creationErr
		}

		return job, nil
	} else {
		return nil, errors.New("nothing found in ERP server")
	}

}

func (jobRepo *JobRepo) Get(jobCode string) (*entity.Job, error) {
	job := entity.Job{}

	getErr := jobRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.CreatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy").
		Preload("Material.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.Factory").
		Preload("UnitOfMeasure.Factory.Address").
		Preload("UnitOfMeasure.Factory.CreatedBy").
		Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure.Factory.UpdatedBy").
		Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.CreatedBy").
		Preload("UnitOfMeasure.CreatedBy.UserRole").
		Preload("UnitOfMeasure.UpdatedBy").
		Preload("UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material").
		Preload("JobItems.Material.UnitOfMeasure").
		Preload("JobItems.Material.UnitOfMeasure.Factory").
		Preload("JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material.CreatedBy").
		Preload("JobItems.Material.CreatedBy.UserRole").
		Preload("JobItems.Material.UpdatedBy").
		Preload("JobItems.Material.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure").
		Preload("JobItems.UnitOfMeasure.Factory").
		Preload("JobItems.UnitOfMeasure.Factory.Address").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.CreatedBy").
		Preload("JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.CreatedBy").
		Preload("JobItems.CreatedBy.UserRole").
		Preload("JobItems.UpdatedBy").
		Preload("JobItems.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItems.JobItemWeighing").
		Preload("JobItems.JobItemWeighing.CreatedBy").
		Preload("JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItems.JobItemWeighing.UpdatedBy").
		Preload("JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("job_code = ?", jobCode).Take(&job).Error
	if getErr != nil {
		return nil, getErr
	}

	return &job, nil
}

func (jobRepo *JobRepo) List(conditions string) ([]entity.Job, error) {
	jobs := []entity.Job{}

	getErr := jobRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.CreatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy").
		Preload("Material.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.Factory").
		Preload("UnitOfMeasure.Factory.Address").
		Preload("UnitOfMeasure.Factory.CreatedBy").
		Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure.Factory.UpdatedBy").
		Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.CreatedBy").
		Preload("UnitOfMeasure.CreatedBy.UserRole").
		Preload("UnitOfMeasure.UpdatedBy").
		Preload("UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material").
		Preload("JobItems.Material.UnitOfMeasure").
		Preload("JobItems.Material.UnitOfMeasure.Factory").
		Preload("JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material.CreatedBy").
		Preload("JobItems.Material.CreatedBy.UserRole").
		Preload("JobItems.Material.UpdatedBy").
		Preload("JobItems.Material.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure").
		Preload("JobItems.UnitOfMeasure.Factory").
		Preload("JobItems.UnitOfMeasure.Factory.Address").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.CreatedBy").
		Preload("JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.CreatedBy").
		Preload("JobItems.CreatedBy.UserRole").
		Preload("JobItems.UpdatedBy").
		Preload("JobItems.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItems.JobItemWeighing").
		Preload("JobItems.JobItemWeighing.CreatedBy").
		Preload("JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItems.JobItemWeighing.UpdatedBy").
		Preload("JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&jobs).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobs, nil
}

func (jobRepo *JobRepo) Update(id string, update *entity.Job) (*entity.Job, error) {
	existingJob := entity.Job{}
	getErr := jobRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingJob).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobRepo.DB.Table(entity.Job{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Job{}
	jobRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
