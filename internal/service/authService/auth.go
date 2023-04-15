package authService

import (
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
	"github.com/hifat/hifat-blog-api/internal/utils/ernos"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{userRepo}
}

func (u authService) Register(req domain.RequestRegister) (res *domain.ResponseRegister, err error) {
	exists, err := u.userRepo.CheckExists("username", req.Username, nil)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ernos.HasAlreadyExists("username")
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	return u.userRepo.Register(req)
}
