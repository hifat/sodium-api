package authService

import (
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{userRepo}
}

func (u authService) Register(req domain.PayloadUser) (res *domain.ResponseUser, validateErors utils.ValidatorType, err error) {
	validateErors, err = utils.Validator(req)
	if err != nil || len(validateErors) > 0 {
		return nil, validateErors, err
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, nil, err
	}

	res, err = u.userRepo.Create(req)

	return res, validateErors, err
}
