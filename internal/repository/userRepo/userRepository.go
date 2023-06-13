package userRepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"gorm.io/gorm"
)

var UserRepoSet = wire.NewSet(NewUserRepository)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db}
}

func (r userRepository) GetFieldsByID(ctx context.Context, ID uuid.UUID, field string) (value interface{}, err error) {
	var fields []interface{}
	return fields[0], r.db.Model(&gormModel.User{}).
		Select(field).
		Where("id = ?", ID).
		Pluck(field, &fields).Error
}
