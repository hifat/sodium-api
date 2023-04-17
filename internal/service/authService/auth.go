package authService

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/constants"
	"github.com/hifat/sodium-api/internal/domain"
	"github.com/hifat/sodium-api/internal/utils"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/token"
)

type authService struct {
	authRepo domain.AuthRepository
}

func NewAuthService(authRepo domain.AuthRepository) domain.AuthService {
	return &authService{authRepo}
}

func (u authService) Register(req domain.RequestRegister, res *domain.ResponseRegister) (err error) {
	exists, err := u.authRepo.CheckUserExists("username", req.Username, nil)
	if err != nil {
		return err
	}

	if exists {
		return ernos.HasAlreadyExists("username")
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	return u.authRepo.Register(req, res)
}

func (u authService) Login(req domain.RequestLogin, res *domain.ResponseLogin) (err error) {
	var user domain.ResponseLoginRepo
	err = u.authRepo.Login(req, &user)
	if err != nil {
		if err.Error() == ernos.M.RECORD_NOTFOUND {
			return ernos.Other(ernos.Ernos{
				Message: "Username or password is incorrect",
				Code:    ernos.C.INVALID_CREDENTIALS,
			})
		}

		return err
	}

	userPayload := token.UserPayload{
		UserID:   user.ID,
		Username: user.Username,
	}

	expired := time.Now().AddDate(0, 0, 7)

	accessSecret := os.Getenv(constants.ACCESS_TOKEN_SECRET)
	accessToken, _, err := token.CreateToken(accessSecret, userPayload, time.Minute*15)
	if err != nil {
		log.Println(err.Error())
		return ernos.InternalServerError("")
	}

	refreshSecret := os.Getenv(constants.REFRESH_TOKEN_SECRET)
	refreshToken, _, err := token.CreateToken(refreshSecret, userPayload, time.Until(expired))
	if err != nil {
		log.Println(err.Error())
		return ernos.InternalServerError("")
	}

	*res = domain.ResponseLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

func (u authService) Logout(ID uuid.UUID) (err error) {
	return nil
}
