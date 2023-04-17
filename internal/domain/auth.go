package domain

import "github.com/google/uuid"

type AuthService interface {
	Register(req RequestRegister, res *ResponseRegister) (err error)
	Login(req RequestLogin, res *ResponseLogin) (err error)
	Logout(ID uuid.UUID) (err error)
}

type AuthRepository interface {
	CheckUserExists(col, value string, exceptID *any) (exists bool, err error)
	Register(req RequestRegister, res *ResponseRegister) (err error)
	Login(req RequestLogin, res *ResponseLoginRepo) (err error)
	Logout(ID uuid.UUID) (err error)
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
	Username string `binding:"required,max=100" json:"username"`
	Password string `binding:"required,max=100" json:"password"`
}

type ResponseLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ResponseLoginRepo struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}
