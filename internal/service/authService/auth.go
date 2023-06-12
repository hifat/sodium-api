package authService

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/constants"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/utils"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/token"
)

var AuthServiceSet = wire.NewSet(NewAuthService)

type authService struct {
	authRepo authDomain.AuthRepository
	userRepo userDomain.UserRepository
}

func NewAuthService(authRepo authDomain.AuthRepository, userRepo userDomain.UserRepository) authDomain.AuthService {
	return &authService{
		authRepo,
		userRepo,
	}
}

func (u authService) Register(req authDomain.RequestRegister, res *authDomain.ResponseRegister) (err error) {
	exists, err := u.authRepo.CheckUserExists("username", req.Username, nil)
	if err != nil {
		log.Println(err)
		return ernos.InternalServerError()
	}

	if exists {
		return ernos.HasAlreadyExists("username")
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return ernos.InternalServerError()
	}

	return u.authRepo.Register(req, res)
}

func (u authService) Login(req authDomain.RequestLogin, res *authDomain.ResponseRefreshToken) (err error) {
	var user authDomain.ResponseRefreshTokenRepo
	err = u.authRepo.Login(req, &user)
	if err != nil {
		if err.Error() == ernos.M.RECORD_NOTFOUND {
			return ernos.Other(ernos.Ernos{
				Status:  http.StatusUnauthorized,
				Message: "Username or password is incorrect",
				Code:    ernos.C.INVALID_CREDENTIALS,
			})
		}

		return err
	}

	newRefreshToken := authDomain.RequestCreateRefreshToken{
		Agent:    req.Agent,
		ClientIP: req.ClientIP,
		UserID:   user.ID,
		Username: user.Username,
	}

	newToken, err := u.CreateRefreshToken(newRefreshToken)
	if err != nil {
		log.Println(err.Error())
		return ernos.InternalServerError()
	}

	*res = authDomain.ResponseRefreshToken{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
	}

	return nil
}

func (u authService) Logout(refreshTokenID uuid.UUID) (err error) {
	return u.authRepo.Logout(refreshTokenID)
}

func (u authService) CreateRefreshToken(req authDomain.RequestCreateRefreshToken) (res *authDomain.ResponseRefreshToken, err error) {
	userPayload := token.UserPayload{
		UserID:   req.UserID,
		Username: fmt.Sprintf("%v", req.Username),
	}

	accessSecret := os.Getenv(constants.ACCESS_TOKEN_SECRET)
	accessToken, _, err := token.CreateToken(accessSecret, userPayload, time.Minute*15)
	if err != nil {
		log.Println(err.Error())
		return nil, ernos.InternalServerError()
	}

	expired := time.Now().AddDate(0, 0, 7)
	refreshSecret := os.Getenv(constants.REFRESH_TOKEN_SECRET)
	refreshToken, _, err := token.CreateToken(refreshSecret, userPayload, time.Until(expired))
	if err != nil {
		log.Println(err.Error())
		return nil, ernos.InternalServerError()
	}
	newRefreshToken := authDomain.RequestCreateRefreshToken{
		Token:    refreshToken,
		Agent:    req.Agent,
		ClientIP: req.ClientIP,
		UserID:   req.UserID,
	}

	_, err = u.authRepo.CreateRefreshToken(newRefreshToken)
	if err != nil {
		log.Println(err.Error())
		return nil, ernos.InternalServerError()
	}

	return &authDomain.ResponseRefreshToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
