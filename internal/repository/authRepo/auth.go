package authRepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/model/gormModel"
	"github.com/hifat/sodium-api/internal/utils/utime"
	"gorm.io/gorm"
)

var AuthRepoSet = wire.NewSet(NewAuthRepository)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authDomain.AuthRepository {
	return &authRepository{db}
}

func (r authRepository) Register(ctx context.Context, req authDomain.RequestRegister, res *authDomain.ResponseRegister) (err error) {
	newUser := gormModel.User{
		Username:  req.Username,
		Password:  req.Password,
		Name:      req.Name,
		CreatedAt: utime.Now(),
		UpdatedAt: utime.Now(),
	}

	err = r.db.Create(&newUser).Error
	if err != nil {
		return err
	}

	return r.db.Model(&gormModel.User{}).Where("id = ?", newUser.ID).First(&res).Error
}

func (r authRepository) Login(ctx context.Context, req authDomain.RequestLogin, res *authDomain.ResponseRefreshTokenRepo) (err error) {
	return r.db.Model(&gormModel.User{}).
		Select("id", "username", "password", "name").
		Where("username = ?", req.Username).
		First(&res).Error
}

func (r authRepository) Logout(ctx context.Context, refreshTokenID uuid.UUID) (err error) {
	return r.db.Where("id = ?", refreshTokenID).
		Delete(&gormModel.RefreshToken{}).Error
}

func (r authRepository) CreateRefreshToken(ctx context.Context, req authDomain.RequestCreateRefreshToken) (res *authDomain.ResponseCreateRefreshToken, err error) {
	refreshToken := gormModel.RefreshToken{
		ID:        req.ID,
		Token:     req.Token,
		Agent:     req.Agent,
		ClientIP:  req.ClientIP,
		UserID:    req.UserID,
		CreatedAt: utime.Now(),
		UpdatedAt: utime.Now(),
	}

	return res, r.db.Create(&refreshToken).
		Scan(&res).Error
}

func (r authRepository) GetRefreshTokenByID(ctx context.Context, refreshTokenID uuid.UUID, res *authDomain.ResponseRefreshTokenClaim) (err error) {
	return r.db.Model(&gormModel.RefreshToken{}).
		Joins("User").
		Where("refresh_tokens.id = ?", refreshTokenID).
		First(&res).Error
}

func (r authRepository) CountLogin(ctx context.Context, col, value string, amount *int64) (err error) {
	return r.db.Model(&gormModel.RefreshToken{}).
		Where(col+" = ?", value).
		Where("is_active IS TRUE").
		Count(amount).Error
}
