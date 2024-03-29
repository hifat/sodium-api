package authDomain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
)

type AuthService interface {
	Register(ctx context.Context, req RequestRegister, res *ResponseRegister) (err error)
	Login(ctx context.Context, req RequestLogin, res *ResponseRefreshToken) (err error)
	Logout(ctx context.Context, refreshTokenID uuid.UUID) (err error)
	CreateRefreshToken(ctx context.Context, req RequestCreateRefreshToken) (res *ResponseRefreshToken, err error)
}

//go:generate mockgen -source=./auth.go -destination=../../repository/authRepo/mockAuthRepo/mockAuthRepo.go -package=mockAuthRepo
type AuthRepository interface {
	Register(ctx context.Context, req RequestRegister, res *ResponseRegister) (err error)
	Login(ctx context.Context, req RequestLogin, res *ResponseRefreshTokenRepo) (err error)
	Logout(ctx context.Context, refreshTokenID uuid.UUID) (err error)
	CreateRefreshToken(ctx context.Context, req RequestCreateRefreshToken) (res *ResponseCreateRefreshToken, err error)
	GetRefreshTokenByID(ctx context.Context, refreshTokenID uuid.UUID, res *ResponseRefreshTokenClaim) (err error)
	CountLogin(ctx context.Context, col, value string, amount *int64) (err error)
}

type RequestRegister struct {
	Username string `binding:"required,max=100" json:"username"`
	Password string `binding:"required,min=8,max=100" json:"password"`
	Name     string `binding:"required,max=100" json:"name"`
}

type ResponseRegister struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type RequestLogin struct {
	Username string   `binding:"required,max=100" json:"username"`
	Password string   `binding:"required,max=100" json:"password"`
	Agent    string   `json:"-"`
	ClientIP utype.IP `json:"-"`
}

type ResponseRefreshToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ResponseRefreshTokenRepo struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

type RequestCreateRefreshToken struct {
	ID       uuid.UUID `json:"ID"`
	Token    string    `json:"token"`
	Agent    string    `json:"agent"`
	ClientIP utype.IP  `json:"clientIP" swaggertype:"string"`
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"-"`
}

type ResponseCreateRefreshToken struct {
	Token    string    `json:"token"`
	Agent    string    `json:"agent"`
	ClientIP utype.IP  `json:"clientIP" swaggertype:"string" gorm:"type:inet"`
	UserID   uuid.UUID `json:"userID"`
	IsActive bool      `json:"isActive"`
}

type ResponseRefreshTokenClaim struct {
	ID        uuid.UUID `json:"ID"`
	Token     string    `json:"token"`
	Agent     string    `json:"agent"`
	ClientIP  utype.IP  `swaggertype:"string" json:"clientIP"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	UserID uuid.UUID               `json:"userID"`
	User   userDomain.ResponseUser `json:"user"`
}

type RequestToken struct {
	RefreshToken string `binding:"required" json:"refreshToken"`
}
