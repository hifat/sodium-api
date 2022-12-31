package domain

import (
	"github.com/hifat/hifat-blog-api/internal/utils"
)

type AuthService interface {
	Register(req PayloadUser) (res *ResponseUser, validateErors utils.ValidatorType, err error)
}

type ResponseRegister struct {
}
