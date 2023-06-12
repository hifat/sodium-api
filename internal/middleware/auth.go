package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/domain/middlewareDomain"
	"github.com/hifat/sodium-api/internal/handler/httpResponse"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/validity"
)

var AuthMiddlewareSet = wire.NewSet(NewAuthMiddleware)

type AuthMiddleware struct {
	authMiddlewareService middlewareDomain.AuthMiddlewareService
}

func NewAuthMiddleware(am middlewareDomain.AuthMiddlewareService) AuthMiddleware {
	return AuthMiddleware{authMiddlewareService: am}
}

func (m AuthMiddleware) AuthGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			httpResponse.BadRequest(ctx, errors.New(ernos.M.NO_AUTHORIZATION_HEADER))
			return
		}

		payload, err := m.authMiddlewareService.AuthGuard(authHeader)
		if err != nil {
			httpResponse.Error(ctx, err)
			return
		}

		ctx.Set("credentials", payload)
		ctx.Next()
	}
}

func (m AuthMiddleware) AuthRefreshGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req authDomain.RequestToken
		err := ctx.ShouldBind(&req)
		if err != nil {
			httpResponse.FormErr(ctx, validity.Validate(err))
			return
		}

		payload, err := m.authMiddlewareService.AuthRefreshGuard(req.RefreshToken)
		if err != nil {
			httpResponse.Error(ctx, err)
			return
		}

		ctx.Set("credentials", payload)
		ctx.Next()
	}
}
