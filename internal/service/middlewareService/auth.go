package middlewareService

import (
	"os"
	"strings"

	"github.com/hifat/sodium-api/internal/constants"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/domain/middlewareDomain"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/token"
)

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
			Message: ernos.M.NOT_FOUND_BEARER,
			Code:    ernos.C.NOT_FOUND_BEARER,
		})
	}

	payload, err = token.VerifyToken(os.Getenv(constants.REFRESH_TOKEN_SECRET), accessToken)
	if err != nil {
		return nil, ernos.Other(ernos.Ernos{
			Message: ernos.M.BROKEN_TOKEN,
			Code:    ernos.C.BROKEN_TOKEN,
		})
	}

	return payload, nil
}

func (s authMiddlewareService) AuthRefreshGuard(refreshToken string) (payload *token.Payload, err error) {
	payloadRefresh, err := token.VerifyToken(os.Getenv(constants.REFRESH_TOKEN_SECRET), refreshToken)
	if err != nil {
		return nil, ernos.Unauthorized(ernos.C.BROKEN_TOKEN)
	}

	// Check if the refresh token is active.
	var claim authDomain.ResponseRefreshTokenClaim
	err = s.authRepo.GetRefreshTokenByID(payloadRefresh.ID, &claim)
	if err != nil {
		return nil, err
	}

	if !claim.IsActive {
		return nil, ernos.Unauthorized("")
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
