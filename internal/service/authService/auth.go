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

func (u authService) Register(req domain.FormRegister) (res *domain.ResponseRegister, err error) {
	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	return u.authRepo.Register(req)
}
