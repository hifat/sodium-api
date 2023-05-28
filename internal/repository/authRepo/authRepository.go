package authRepo

import (
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authDomain.AuthRepository {
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

func (r authRepository) Register(req authDomain.RequestRegister, res *authDomain.ResponseRegister) (err error) {
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

func (r authRepository) Login(req authDomain.RequestLogin, res *authDomain.ResponseRefreshTokenRepo) (err error) {
	return r.db.Model(&gormModel.User{}).
		Select("id", "username", "password", "name").
		Where("username = ?", req.Username).
		First(&res).Error
}

func (r authRepository) Logout(refreshTokenID uuid.UUID) (err error) {
	return r.db.Where("id", refreshTokenID).
		Delete(&gormModel.RefreshToken{}).Error
}

func (r authRepository) CreateRefreshToken(req authDomain.RequestCreateRefreshToken) (res *authDomain.ResponseCreateRefreshToken, err error) {
	refreshToken := gormModel.RefreshToken{
		ID:       req.ID,
		Token:    req.Token,
		Agent:    req.Agent,
		ClientIP: req.ClientIP,
		UserID:   req.UserID,
	}

	return res, r.db.Create(&refreshToken).Scan(&res).Error
}

func (r authRepository) GetRefreshTokenByID(refreshTokenID uuid.UUID, res *authDomain.ResponseRefreshTokenClaim) (err error) {
	return r.db.Model(&gormModel.RefreshToken{}).
		Joins("User").
		Where("refresh_tokens.id = ?", refreshTokenID).
		First(&res).Error
}
