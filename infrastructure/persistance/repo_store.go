package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/infrastructure/config"
	"os"

	"github.com/go-redis/redis"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepoStore struct {
	DB     *gorm.DB
	Cache  *redis.Client
	Logger hclog.Logger
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

	repoStore.DB = gormDB
	repoStore.Logger = logging
	repoStore.Cache = cacheStore.RedisClient

	return &repoStore, nil
}

func (repoStore *RepoStore) Migrate() error {
	return repoStore.DB.AutoMigrate(
		&entity.UserRole{},
		&entity.User{},
		&entity.UserRoleAccess{},
		&entity.Address{},
		&entity.Company{},
		&entity.Factory{},
		&entity.UserFactoryAccess{},
		&entity.UnitOfMeasure{},
		&entity.UnitOfMeasureConversion{},
		&entity.Terminal{},
		&entity.Material{},
		&entity.BOM{},
		&entity.BOMItem{},
		&entity.Job{},
		&entity.JobItem{},
		&entity.Shift{},
		&entity.ShiftSchedule{},
		&entity.JobAssignment{},
		&entity.OverIssue{},
		&entity.UnderIssue{},
	)
}
