package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/hifat/hifat-blog-api/internal/utils"
)

/* -------------------------------------------------------------------------- */
/*                                    Auth                                    */
/* -------------------------------------------------------------------------- */

type AuthRepository interface {
	Register(req FormRegister) (res *ResponseRegister, err error)
}

type AuthService interface {
	Register(req FormRegister) (res *ResponseRegister, err error, validateErors utils.ValidatorType)
}

/* -------------------------------------------------------------------------- */
/*                                  Register                                  */
/* -------------------------------------------------------------------------- */

type FormRegister struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Name     string `validate:"required,max=100" json:"name"`
}

type ResponseRegister struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
