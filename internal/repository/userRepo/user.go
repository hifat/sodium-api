package userRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"github.com/hifat/sodium-api/internal/utils/helper/hfunk"
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

func (r userRepository) CheckExists(ctx context.Context, col, expectValue string) (exists bool, err error) {
	columns := []any{
		"username",
		"password",
		"name",
		"created_at",
		"updated_at",
	}
	isIncludesCol := hfunk.Includes(columns, col)

	if !isIncludesCol {
		return false, errors.New("col must be includes " + fmt.Sprintf("%v", columns))
	}

	if expectValue == "" {
		return false, errors.New("expectValue must be required")
	}

	tx := r.db.Model(&gormModel.User{}).
		Select(`COUNT(*) > 0`).
		Where(col+" = ?", expectValue)

	return exists, tx.Find(&exists).Error
}
