package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/domain/middlewareDomain"
	"github.com/hifat/sodium-api/internal/handler/httpResponse"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/validity"
)

type authMiddleware struct {
	authMiddlewareService middlewareDomain.AuthMiddlewareService
}

func NewAuthMiddleware(authMiddlewareService middlewareDomain.AuthMiddlewareService) *authMiddleware {
	return &authMiddleware{authMiddlewareService}
}

func (m authMiddleware) AuthGuard(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader != "" {
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

func (m authMiddleware) AuthRefreshGuard(ctx *gin.Context) {
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
