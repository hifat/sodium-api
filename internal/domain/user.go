package domain

type UserRepository interface {
	Create(req PayloadUser) (res *ResponseUser, err error)
}

type PayloadUser struct {
	Username string `binding:"required,max=100" json:"username"`
	Password string `binding:"required,min=8,max=100" json:"password"`
	Name     string `binding:"required,max=100" json:"name"`
}

type ResponseUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
