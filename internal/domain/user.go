package domain

type UserRepository interface {
	Create(req PayloadUser) (res *ResponseUser, err error)
}

type PayloadUser struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Name     string `validate:"required,max=100" json:"name"`
}

type ResponseUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
