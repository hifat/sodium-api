package middlewareDomain

import "github.com/hifat/sodium-api/internal/utils/token"

type AuthMiddlewareService interface {
	AuthGuard(authTokenHeader string) (payload *token.Payload, err error)
	AuthRefreshGuard(refreshToken string) (payload *token.Payload, err error)
}
