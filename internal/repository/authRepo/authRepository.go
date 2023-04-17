package authRepo

import (
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewauthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepository{db}
}

func (r authRepository) CheckUserExists(col, value string, exceptID *any) (exists bool, err error) {
	tx := r.db.Model(&gormModel.User{}).
		Select(`COUNT(*) > 0`).
		Where("username = ?", value)

	if exceptID != nil {
		tx.Where("id = ?", exceptID)
	}

	err = tx.Find(&exists).Error

	return exists, err
}

func (r authRepository) Register(req domain.RequestRegister, res *domain.ResponseRegister) (err error) {
	newUser := gormModel.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}

	err = r.db.Create(&newUser).Error
	if err != nil {
		return err
	}

	return r.db.Model(&gormModel.User{}).Where("id = ?", newUser.ID).First(&res).Error
}

func (r authRepository) Login(req domain.RequestLogin, res *domain.ResponseLoginRepo) (err error) {
	return r.db.Model(&gormModel.User{}).
		Select("id", "username", "password", "name").
		Where("username = ?", req.Username).
		First(&res).Error
}

func (r authRepository) Logout(ID uuid.UUID) (err error) {
	return nil
}
