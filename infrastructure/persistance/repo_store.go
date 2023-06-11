package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"eazyweigh/infrastructure/config"
	"net/url"
	"os"

	"github.com/go-redis/redis"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepoStore struct {
	DB                          *gorm.DB
	WarehouseDB                 *gorm.DB
	Cache                       *redis.Client
	Logger                      hclog.Logger
	AddressRepo                 repository.AddressRepository
	AuthRepo                    repository.AuthRepository
	BatchRepo                   repository.BatchRepository
	BOMRepo                     repository.BOMRepository
	BOMItemRepo                 repository.BOMItemsRepository
	CommonRepo                  repository.CommonRepository
	CompanyRepo                 repository.CompanyRepository
	DeviceRepo                  repository.DeviceRepository
	DeviceDataRepo              repository.DeviceDataRepository
	DeviceTypeRepo              repository.DeviceTypeRepository
	FactoryRepo                 repository.FactoryRepository
	JobRepo                     repository.JobRepository
	JobItemRepo                 repository.JobItemRepository
	JobItemAssignmentRepo       repository.JobItemAssignmentRepository
	JobItemWeighingRepo         repository.JobItemWeighingRepository
	MaterialRepo                repository.MaterialRepository
	OverIssueRepo               repository.OverIssueRepository
	ProcessRepo                 repository.ProcessRepository
	ScannedDataRepo             repository.ScannedDataRepository
	ShiftRepo                   repository.ShiftRepository
	ShiftScheduleRepo           repository.ShiftScheduleRepository
	StepRepo                    repository.StepRepository
	StepTypeRepo                repository.StepTypeRepository
	TerminalRepo                repository.TerminalRepository
	UnderIssueRepo              repository.UnderIssueRepository
	UnitOfMeasureRepo           repository.UnitOfMeasureRepository
	UnitOfMeasureConversionRepo repository.UnitOfMeasureConversionRepository
	UserRepo                    repository.UserRepository
	UserRoleRepo                repository.UserRoleRepository
	UserRoleAccessRepo          repository.UserRoleAccessRepository
	UserCompanyRepo             repository.UserCompanyRepository
	UserFactoryRepo             repository.UserFactoryRepository
	VesselRepo                  repository.VesselRepository
}

func NewRepoStore(serverConfig *config.ServerConfig, logging hclog.Logger) (*RepoStore, error) {
	repoStore := RepoStore{}
	dbConfig := config.NewDatabaseConfig()

	// Get Caching Service
	cacheStore, cacheError := NewCacheStore(*serverConfig)
	if cacheError != nil {
		logging.Error(cacheError.Error())
		os.Exit(1)
	}

	// MySQL Connection
	mysqlURL := dbConfig.DbUser + ":" + dbConfig.DbPassword + "@tcp(" + dbConfig.DbHost + ":" + dbConfig.DbPort + ")/" + dbConfig.DbName + "?parseTime=True"
	gormDB, gormErr := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	if gormErr != nil {
		return nil, gormErr
	}
	sqlDB, _ := gormDB.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(10000)

	// MSSQL Connection
	username := url.QueryEscape(dbConfig.WarehouseUser)
	password := url.QueryEscape(dbConfig.WarehousePassword)
	sqlURL := "sqlserver://" + username + ":" + password + "@" + dbConfig.WarehouseHost + ":1433?database=" + dbConfig.WarehouseDBName
	remoteSQLDB, _ := gorm.Open(sqlserver.Open(sqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	// if sqlErr != nil {
	// 	return nil, sqlErr
	// }
	repoStore.WarehouseDB = remoteSQLDB

	repoStore.DB = gormDB
	repoStore.Logger = logging
	repoStore.Cache = cacheStore.RedisClient
	repoStore.AddressRepo = NewAddressRepo(gormDB, logging)
	repoStore.AuthRepo = NewAuthRepo(logging, serverConfig, cacheStore.RedisClient)
	repoStore.BatchRepo = NewBatchRepo(gormDB, logging)
	repoStore.BOMItemRepo = NewBOMItemRepo(gormDB, logging)
	repoStore.BOMRepo = NewBOMRepo(gormDB, logging, repoStore.BOMItemRepo)
	repoStore.CompanyRepo = NewCompanyRepo(gormDB, logging)
	repoStore.CommonRepo = NewCommonRepo(gormDB, logging)
	repoStore.DeviceRepo = NewDeviceRepo(gormDB, logging)
	repoStore.DeviceDataRepo = NewDeviceDataRepo(gormDB, logging)
	repoStore.DeviceTypeRepo = NewDeviceTypeRepo(gormDB, logging)
	repoStore.FactoryRepo = NewFactoryRepo(gormDB, logging)
	repoStore.JobRepo = NewJobRepo(gormDB, remoteSQLDB, logging)
	repoStore.JobItemRepo = NewJobItemRepo(gormDB, logging)
	repoStore.JobItemAssignmentRepo = NewJobItemAssignmentRepo(gormDB, logging)
	repoStore.JobItemWeighingRepo = NewJobItemWeighingRepo(gormDB, logging)
	repoStore.MaterialRepo = NewMaterialRepo(gormDB, logging)
	repoStore.OverIssueRepo = NewOverIssueRepo(gormDB, logging)
	repoStore.ProcessRepo = NewProcessRepo(gormDB, logging)
	repoStore.ScannedDataRepo = NewScannedDataRepo(gormDB, logging)
	repoStore.ShiftRepo = NewShiftRepo(gormDB, logging)
	repoStore.ShiftScheduleRepo = NewShiftScheduleRepo(gormDB, logging)
	repoStore.StepRepo = NewStepRepo(gormDB, logging)
	repoStore.StepTypeRepo = NewStepTypeRepo(gormDB, logging)
	repoStore.TerminalRepo = NewTerminalRepo(gormDB, logging)
	repoStore.UnderIssueRepo = NewUnderIssueRepo(gormDB, logging)
	repoStore.UnitOfMeasureRepo = NewUnitOfMeasureRepo(gormDB, logging)
	repoStore.UnitOfMeasureConversionRepo = NewUnitOfMeasureConversionRepo(gormDB, logging)
	repoStore.UserRepo = NewUserRepo(gormDB, logging)
	repoStore.UserRoleRepo = NewUserRoleRepo(gormDB, logging)
	repoStore.UserRoleAccessRepo = NewUserRoleAccessRepo(gormDB, logging)
	repoStore.UserCompanyRepo = NewUserCompanyRepo(gormDB, logging)
	repoStore.UserFactoryRepo = NewUserFactoryRepo(gormDB, logging)
	repoStore.VesselRepo = NewVesselRepo(gormDB, logging)

	return &repoStore, nil
}

func (repoStore *RepoStore) Migrate() error {
	return repoStore.DB.AutoMigrate(
		&entity.UserRole{},
		&entity.User{},
		&entity.Address{},
		&entity.Company{},
		&entity.Factory{},
		&entity.UserCompany{},
		&entity.UserFactory{},
		&entity.UserRoleAccess{},
		&entity.UserFactoryAccess{},
		&entity.UserCompanyAccess{},
		&entity.UnitOfMeasure{},
		&entity.UnitOfMeasureConversion{},
		&entity.Terminal{},
		&entity.UserTerminalAccess{},
		&entity.Material{},
		&entity.BOM{},
		&entity.BOMItem{},
		&entity.Job{},
		&entity.JobItem{},
		&entity.Shift{},
		&entity.ShiftSchedule{},
		&entity.JobItemAssignment{},
		&entity.JobItemWeighing{},
		&entity.OverIssue{},
		&entity.UnderIssue{},
		&entity.ScannedData{},
		&entity.StepType{},
		&entity.Process{},
		&entity.Step{},
		&entity.Vessel{},
		&entity.DeviceType{},
		&entity.Device{},
		&entity.DeviceData{},
		&entity.Batch{},
	)
}
