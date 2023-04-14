package domain

type AuthService interface {
	Register(req RequestRegister) (res *ResponseRegister, err error)
}
