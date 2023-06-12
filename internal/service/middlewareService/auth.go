package middlewareService

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/constants"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/domain/middlewareDomain"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/token"
)

var AuthMiddlewareServiceSet = wire.NewSet(NewAuthMiddlewareService)

type authMiddlewareService struct {
	authRepo authDomain.AuthRepository
}

func NewAuthMiddlewareService(authRepo authDomain.AuthRepository) middlewareDomain.AuthMiddlewareService {
	return &authMiddlewareService{authRepo}
}

func (s authMiddlewareService) AuthGuard(authTokenHeader string) (payload *token.Payload, err error) {
	accessToken := strings.TrimPrefix(authTokenHeader, "Bearer ")
	if accessToken == authTokenHeader {
		return nil, ernos.Other(ernos.Ernos{
			Status:  http.StatusBadRequest,
			Message: ernos.M.NOT_FOUND_BEARER,
			Code:    ernos.C.NOT_FOUND_BEARER,
		})
	}

	payload, err = token.VerifyToken(os.Getenv(constants.ACCESS_TOKEN_SECRET), accessToken)
	if err != nil {
		return nil, ernos.Other(ernos.Ernos{
			Status:  http.StatusUnauthorized,
			Message: ernos.M.BROKEN_TOKEN,
			Code:    ernos.C.BROKEN_TOKEN,
		})
	}

	return payload, nil
}

func (s authMiddlewareService) AuthRefreshGuard(refreshToken string) (payload *token.Payload, err error) {
	payloadRefresh, err := token.VerifyToken(os.Getenv(constants.REFRESH_TOKEN_SECRET), refreshToken)
	if err != nil {
		log.Println(err.Error())
		return nil, ernos.Unauthorized(ernos.C.BROKEN_TOKEN)
	}

	// Check if the refresh token is active.
	var claim authDomain.ResponseRefreshTokenClaim
	err = s.authRepo.GetRefreshTokenByID(payloadRefresh.ID, &claim)
	if err != nil {
		// TODO reflect check record not found
		if err.Error() == ernos.M.RECORD_NOTFOUND {
			return nil, ernos.NotFound("refresh token")
		}

		log.Println(err.Error())
		return nil, ernos.InternalServerError()
	}

	if !claim.IsActive {
		log.Println(err.Error())
		return nil, ernos.Unauthorized()
	}

	payload = &token.Payload{
		ID:        claim.ID,
		UserID:    claim.UserID,
		Username:  payloadRefresh.Username,
		IssuedAt:  payloadRefresh.IssuedAt,
		ExpiredAt: payloadRefresh.ExpiredAt,
	}

	return payload, nil
}
