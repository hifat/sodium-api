package userRepo

import (
	"github.com/hifat/sodium-api/internal/domain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r userRepository) CheckExists(col, value string, exceptID *any) (exists bool, err error) {
	tx := r.db.Model(&gormModel.User{}).
		Select(`COUNT(*) > 0`).
		Where("username = ?", value)

	if exceptID != nil {
		tx.Where("id = ?", exceptID)
	}

	err = tx.Find(&exists).Error

	return exists, err
}

func (r userRepository) Register(req domain.RequestRegister) (res *domain.ResponseRegister, err error) {
	newUser := gormModel.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}

	err = r.db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return res, r.db.Model(&gormModel.User{}).Where("id = ?", newUser.ID).First(&res).Error
}
