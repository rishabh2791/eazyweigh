package persistance

import (
	"eazyweigh/domain/entity"
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
	AddressRepo                 *AddressRepo
	AuthRepo                    *AuthRepo
	BOMRepo                     *BOMRepo
	BOMItemRepo                 *BOMItemRepo
	CommonRepo                  *CommonRepo
	CompanyRepo                 *CompanyRepo
	FactoryRepo                 *FactoryRepo
	JobRepo                     *JobRepo
	JobItemRepo                 *JobItemRepo
	JobItemAssignmentRepo       *JobItemAssignmentRepo
	JobItemWeighingRepo         *JobItemWeighingRepo
	MaterialRepo                *MaterialRepo
	OverIssueRepo               *OverIssueRepo
	ScannedDataRepo             *ScannedDataRepo
	ShiftRepo                   *ShiftRepo
	ShiftScheduleRepo           *ShiftScheduleRepo
	TerminalRepo                *TerminalRepo
	UnderIssueRepo              *UnderIssueRepo
	UnitOfMeasureRepo           *UnitOfMeasureRepo
	UnitOfMeasureConversionRepo *UnitOfMeasureConversionRepo
	UserRepo                    *UserRepo
	UserRoleRepo                *UserRoleRepo
	UserRoleAccessRepo          *UserRoleAccessRepo
	UserCompanyRepo             *UserCompanyRepo
	UserFactoryRepo             *UserFactoryRepo
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
	remoteSQLDB, sqlErr := gorm.Open(sqlserver.Open(sqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	if sqlErr != nil {
		return nil, sqlErr
	}
	repoStore.WarehouseDB = remoteSQLDB

	repoStore.DB = gormDB
	repoStore.Logger = logging
	repoStore.Cache = cacheStore.RedisClient
	repoStore.AddressRepo = NewAddressRepo(gormDB, logging)
	repoStore.AuthRepo = NewAuthRepo(logging, serverConfig, cacheStore.RedisClient)
	repoStore.BOMItemRepo = NewBOMItemRepo(gormDB, logging)
	repoStore.BOMRepo = NewBOMRepo(gormDB, logging, repoStore.BOMItemRepo)
	repoStore.CompanyRepo = NewCompanyRepo(gormDB, logging)
	repoStore.CommonRepo = NewCommonRepo(gormDB, logging)
	repoStore.FactoryRepo = NewFactoryRepo(gormDB, logging)
	repoStore.JobRepo = NewJobRepo(gormDB, remoteSQLDB, logging)
	repoStore.JobItemRepo = NewJobItemRepo(gormDB, logging)
	repoStore.JobItemAssignmentRepo = NewJobItemAssignmentRepo(gormDB, logging)
	repoStore.JobItemWeighingRepo = NewJobItemWeighingRepo(gormDB, logging)
	repoStore.MaterialRepo = NewMaterialRepo(gormDB, logging)
	repoStore.OverIssueRepo = NewOverIssueRepo(gormDB, logging)
	repoStore.ScannedDataRepo = NewScannedDataRepo(gormDB, logging)
	repoStore.ShiftRepo = NewShiftRepo(gormDB, logging)
	repoStore.ShiftScheduleRepo = NewShiftScheduleRepo(gormDB, logging)
	repoStore.TerminalRepo = NewTerminalRepo(gormDB, logging)
	repoStore.UnderIssueRepo = NewUnderIssueRepo(gormDB, logging)
	repoStore.UnitOfMeasureRepo = NewUnitOfMeasureRepo(gormDB, logging)
	repoStore.UnitOfMeasureConversionRepo = NewUnitOfMeasureConversionRepo(gormDB, logging)
	repoStore.UserRepo = NewUserRepo(gormDB, logging)
	repoStore.UserRoleRepo = NewUserRoleRepo(gormDB, logging)
	repoStore.UserRoleAccessRepo = NewUserRoleAccessRepo(gormDB, logging)
	repoStore.UserCompanyRepo = NewUserCompanyRepo(gormDB, logging)
	repoStore.UserFactoryRepo = NewUserFactoryRepo(gormDB, logging)

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
	)
}
