package domain

type UserRepository interface {
	Register(req RequestRegister) (res *ResponseRegister, err error)
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
