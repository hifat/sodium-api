package middlewareDomain

import (
	"context"

	"github.com/hifat/sodium-api/internal/utils/token"
)

type AuthMiddlewareService interface {
	AuthGuard(ctx context.Context, authTokenHeader string) (payload *token.Payload, err error)
	AuthRefreshGuard(ctx context.Context, refreshToken string) (payload *token.Payload, err error)
}
