package authService

import (
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
)

type authService struct {
	authRepo domain.AuthRepository
}

func NewAuthService(authRepo domain.AuthRepository) domain.AuthService {
	return &authService{authRepo}
}

func (u authService) Register(req domain.FormRegister) (res *domain.ResponseRegister, validateErors utils.ValidatorType, err error) {
	validateErors, err = utils.Validator(req)
	if err != nil || len(validateErors) > 0 {
		return nil, validateErors, err
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, nil, err
	}

	res, err = u.authRepo.Register(req)

	return res, validateErors, err
}
