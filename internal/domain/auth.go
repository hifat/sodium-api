package domain

type AuthService interface {
	Register(req PayloadUser) (res *ResponseUser, err error)
}

type ResponseRegister struct {
}
