package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UserRepository = &UserRepo{}

func NewUserRepo(db *gorm.DB, logging hclog.Logger) *UserRepo {
	return &UserRepo{
		DB:     db,
		Logger: logging,
	}
}

func (userRepo *UserRepo) Create(user *entity.User, action string) (*entity.User, error) {
	validationErr := user.Validate(action)
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := userRepo.DB.Create(&user).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return user, nil
}

func (userRepo *UserRepo) Get(username string) (*entity.User, error) {
	user := entity.User{}
	getErr := userRepo.DB.Preload(clause.Associations).Where("username = ?", username).Take(&user).Error
	if getErr != nil {
		return nil, getErr
	}
	return &user, nil
}

func (userRepo *UserRepo) List(conditions string) ([]entity.User, error) {
	users := []entity.User{}
	getErr := userRepo.DB.
		Preload(clause.Associations).Where(conditions).Find(&users).Error
	if getErr != nil {
		return nil, getErr
	}
	return users, nil
}

func (userRepo *UserRepo) Update(username string, update map[string]interface{}, user string) (*entity.User, error) {
	existingUser := entity.User{}

	err := userRepo.DB.Where("username = ?", username).Take(&existingUser).Error
	if err != nil {
		return nil, err
	}

	updationErr := userRepo.DB.Table(entity.User{}.Tablename()).Where("username = ?", username).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.User{}
	userRepo.DB.Where("username = ?", username).Take(&updated)

	return &updated, nil
}
