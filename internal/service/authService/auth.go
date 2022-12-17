package authService

import "github.com/hifat/hifat-blog-api/internal/domain"

type authService struct {
	authRepo domain.AuthRepository
}

func NewAuthService(authRepo domain.AuthRepository) domain.AuthService {
	return &authService{authRepo}
}

func (u authService) Register(req domain.FormRegister) (res *domain.ResponseRegister, err error) {
	return u.authRepo.Register(req)
}
