package userRepo

import (
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/model/gormModel"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
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
