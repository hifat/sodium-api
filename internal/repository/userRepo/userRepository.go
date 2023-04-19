package userRepo

import (
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db}
}

func (r userRepository) GetFieldsByID(ID uuid.UUID, field string) (value []interface{}, err error) {
	return value, r.db.Model(&gormModel.User{}).
		Select(field).
		Where("id = ?", ID).
		Pluck(field, &value).Error
}
