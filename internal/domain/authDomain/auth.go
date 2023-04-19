package authDomain

import (
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
)

type AuthService interface {
	Register(req RequestRegister, res *ResponseRegister) (err error)
	Login(req RequestLogin, res *ResponseRefreshToken) (err error)
	Logout(ID uuid.UUID) (err error)
	CreateRefreshToken(req RequestCreateRefreshToken) (res *ResponseRefreshToken, err error)
}

type AuthRepository interface {
	CheckUserExists(col, value string, exceptID *any) (exists bool, err error)
	Register(req RequestRegister, res *ResponseRegister) (err error)
	Login(req RequestLogin, res *ResponseRefreshTokenRepo) (err error)
	Logout(ID uuid.UUID) (err error)
	CreateRefreshToken(req RequestCreateRefreshToken) (res *ResponseCreateRefreshToken, err error)
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
	Token    string    `json:"token"`
	Agent    string    `json:"agent"`
	ClientIP utype.IP  `json:"clientIP"`
	UserID   uuid.UUID `json:"userID"`
}

type ResponseCreateRefreshToken struct {
	Token    string    `json:"token"`
	Agent    string    `json:"agent"`
	ClientIP utype.IP  `json:"clientIP" gorm:"type:inet"`
	UserID   uuid.UUID `json:"userID"`
	IsActive bool      `json:"isActive"`
}
