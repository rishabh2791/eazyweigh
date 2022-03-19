package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TerminalRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.TerminalRepository = &TerminalRepo{}

func NewTerminalRepo(db *gorm.DB, logger hclog.Logger) *TerminalRepo {
	return &TerminalRepo{
		DB:     db,
		Logger: logger,
	}
}

func (terminalRepo *TerminalRepo) Create(termial *entity.Terminal) (*entity.Terminal, error) {
	validationErr := termial.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := terminalRepo.DB.Create(&termial).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return termial, nil
}

func (terminalRepo *TerminalRepo) Get(id string) (*entity.Terminal, error) {
	terminal := entity.Terminal{}
	getErr := terminalRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&terminal).Error
	if getErr != nil {
		return nil, getErr
	}

	return &terminal, nil
}

func (terminalRepo *TerminalRepo) List(conditions string) ([]entity.Terminal, error) {
	terminals := []entity.Terminal{}
	getErr := terminalRepo.DB.Preload(clause.Associations).Where(conditions).Find(&terminals).Error
	if getErr != nil {
		return nil, getErr
	}

	return terminals, nil
}

func (terminalRepo *TerminalRepo) Update(id string, update *entity.Terminal) (*entity.Terminal, error) {
	existingTerminal := entity.Terminal{}
	getErr := terminalRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingTerminal).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := terminalRepo.DB.Table(entity.Terminal{}.Tablename()).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Terminal{}
	terminalRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
