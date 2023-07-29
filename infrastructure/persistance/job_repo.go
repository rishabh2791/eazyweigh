package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"errors"
	"strings"

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

type RemoteMaterial struct {
	StockCode   string
	Description string
}

type RemoteBOMItem struct {
	StockCode string
	Quantity  float32
}

type RemoteJob struct {
	Job       string
	StockCode string
	QtyToMake float32
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

func (jobRepo *JobRepo) GetOpenJobs() ([]RemoteJob, error) {
	remoteJobs := []RemoteJob{}
	jobQuery := "SELECT Job, StockCode, QtyToMake FROM dbo.WipMaster WHERE Complete = 'N' AND StockCode LIKE '70%'"
	rows, getErr := jobRepo.WarehouseDB.Raw(jobQuery).Rows()
	defer rows.Close()
	if getErr != nil {
		return nil, getErr
	}
	for rows.Next() {
		remoteJob := RemoteJob{}
		scanErr := rows.Scan(&remoteJob.Job, &remoteJob.StockCode, &remoteJob.QtyToMake)
		if scanErr != nil {
			return nil, scanErr
		}
		remoteJobs = append(remoteJobs, remoteJob)
	}
	return remoteJobs, nil
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
		bomItem.UpperTolerance = 0.2
		bomItem.LowerTolerance = 0.2
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

func checkInRemoteBOMItems(bomItem *entity.BOMItem, bomItems []RemoteBOMItem) bool {
	for _, remoteBOMItem := range bomItems {
		if bomItem.Material.Code == remoteBOMItem.StockCode && bomItem.Quantity == remoteBOMItem.Quantity {
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
		check := checkInRemoteBOMItems(&bomItem, remoteBOMItems)
		revision = revision && check
	}
	return revision, nil
}

func (jobRepo *JobRepo) GetExistingBOM(stockCode string, boms []entity.BOM) *entity.BOM {
	remoteBOMItems, remoteErr := jobRepo.getBOMFromRemote(stockCode)
	if remoteErr != nil {
		return nil
	}
	for _, bom := range boms {
		checkThisBOM := true
		for _, bomItem := range bom.BOMItems {
			check := checkInRemoteBOMItems(&bomItem, remoteBOMItems)
			checkThisBOM = checkThisBOM && check
		}
		if checkThisBOM {
			return &bom
		}
	}
	return nil
}

func (jobRepo *JobRepo) Create(job *entity.Job) (*entity.Job, error) {
	var stockCode string
	var quantity float32

	stockCode = job.Material.Code
	quantity = job.Quantity

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
			// log.Println(remoteMaterial)

			//Create Material
			material, getErr := jobRepo.CreateMaterial(remoteMaterial, job, "Bulk")
			if getErr != nil {
				return nil, getErr
			}
			existingStockCode = *material
		}
		//Check if BOM Exists, If exists BOM Items are already created. There may be revisions
		existingBOMs := []entity.BOM{}
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
			Preload(clause.Associations).Where("factory_id = ? AND material_id = ?", job.FactoryID, existingStockCode.ID).Find(&existingBOMs).Error
		if getBomErr != nil {
			//Not Created
			bom, creationErr := jobRepo.CreateBOM(job, &existingStockCode, stockCode, 1)
			if creationErr != nil {
				return nil, creationErr
			}
			existingBOM = *bom
		} else {
			//BOM Already Created so check for revision
			existing := jobRepo.GetExistingBOM(stockCode, existingBOMs)
			if existing == nil {
				bom, creationErr := jobRepo.CreateBOM(job, &existingStockCode, stockCode, len(existingBOMs)+1)
				if creationErr != nil {
					return nil, creationErr
				}
				existingBOM = *bom
			} else {
				existingBOM = *existing
			}
		}

		job.Material = nil
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
		return nil, errors.New("nothing found in ERP server\n")
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
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItems").
		Preload("JobItems.Material").
		Preload("JobItems.Material.UnitOfMeasure").
		Preload("JobItems.Material.UnitOfMeasure.Factory").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItems.UnitOfMeasure").
		Preload("JobItems.UnitOfMeasure.Factory").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.CreatedBy").
		Preload("JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.Factory.Address").
		Preload("JobItems.CreatedBy").
		Preload("JobItems.UpdatedBy").
		Preload("JobItems.CreatedBy.UserRole").
		Preload("JobItems.UpdatedBy.UserRole").
		Preload("JobItems.Material.CreatedBy").
		Preload("JobItems.Material.CreatedBy.UserRole").
		Preload("JobItems.Material.UpdatedBy").
		Preload("JobItems.Material.UpdatedBy.UserRole").
		Preload("JobItems.JobItemWeighing").
		Preload("JobItems.JobItemWeighing.CreatedBy").
		Preload("JobItems.JobItemWeighing.UpdatedBy").
		Preload("JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&jobs).Error
	if getErr != nil {
		return nil, getErr
	}
	return jobs, nil
}

func (jobRepo *JobRepo) Update(jobCode string, update *entity.Job) (*entity.Job, error) {
	// Check Existing Job
	existingJob := entity.Job{}
	getErr := jobRepo.DB.Preload(clause.Associations).Where("job_code = ?", jobCode).Take(&existingJob).Error
	if getErr != nil {
		return nil, getErr
	}

	if existingJob.Complete {
		return nil, errors.New("Job " + jobCode + " Complete and can not be updated.\n")
	}

	factor := update.Quantity / existingJob.Quantity

	// Get Job Items for the existing Job
	existingJobItems := []entity.JobItem{}
	weighedJobItems := []entity.JobItem{}
	jobItemsError := jobRepo.DB.Preload(clause.Associations).Where("job_id=?", existingJob.ID).Find(&existingJobItems).Error
	if jobItemsError != nil {
		return nil, jobItemsError
	}
	completedJobItemsError := jobRepo.DB.Preload(clause.Associations).Where("job_id=? AND actual_weight != 0", existingJob.ID).Find(&weighedJobItems).Error
	if completedJobItemsError != nil {
		return nil, completedJobItemsError
	}

	if len(weighedJobItems) != 0 {
		return nil, errors.New("At least one Job Item Weighed Out, can not update Job " + jobCode + ".\n")
	}

	for _, jobItem := range existingJobItems {
		jobItemUpdate := entity.JobItem{}
		jobItemUpdate.RequiredWeight = jobItem.RequiredWeight * factor
		jobItemUpdate.UpperBound = jobItem.UpperBound * factor
		jobItemUpdate.LowerBound = jobItem.LowerBound * factor
		jobItemUpdate.UpdatedByUsername = update.UpdatedByUsername
		jobRepo.DB.Table(entity.JobItem{}.Tablename()).Where("id = ?", jobItem.ID).Updates(jobItemUpdate)
	}

	// Update Job Details
	updationErr := jobRepo.DB.Table(entity.Job{}.Tablename()).Where("job_code = ?", jobCode).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Job{}
	jobRepo.DB.
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
		Preload(clause.Associations).Where("job_code = ?", jobCode).Take(&updated)

	return &updated, nil
}

func (jobRepo *JobRepo) PullFromRemote(factoryID string, username string) error {
	remoteJobs, remoteErr := jobRepo.GetOpenJobs()
	error := ""
	if remoteErr != nil {
		return remoteErr
	}
	unitOfMeasure := entity.UnitOfMeasure{}
	getErr := jobRepo.DB.Where("factory_id LIKE ? AND code LIKE ?", factoryID, "KG").Take(&unitOfMeasure).Error
	if getErr != nil {
		return getErr
	}
	for _, remoteJob := range remoteJobs {
		job := entity.Job{}
		job.FactoryID = factoryID
		job.JobCode = remoteJob.Job[9:len(remoteJob.Job)]
		job.Material = &entity.Material{}
		job.Material.FactoryID = factoryID
		job.Material.Code = remoteJob.StockCode
		job.Material.UnitOfMeasureID = unitOfMeasure.ID
		job.Quantity = remoteJob.QtyToMake
		job.UnitOfMeasureID = unitOfMeasure.ID
		job.CreatedByUsername = username
		job.UpdatedByUsername = username
		_, jobCreationError := jobRepo.Create(&job)
		if jobCreationError != nil {
			if strings.Contains(jobCreationError.Error(), "Duplicate") {
				existingJob := entity.Job{}
				jobRepo.DB.Where("job_code = ?", job.JobCode).Take(&existingJob)
				if job.Quantity != existingJob.Quantity {
					updatedJob := entity.Job{}
					updatedJob.Quantity = job.Quantity
					_, updationErr := jobRepo.Update(job.JobCode, &updatedJob)
					if updationErr != nil {
						error += updationErr.Error()
					}
				} else {
					error += "Job " + job.JobCode + " already created.\n"
				}
			} else {
				error += jobCreationError.Error() + "\n"
			}
		}
	}
	return errors.New(error)
}
