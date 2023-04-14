package authHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
	"github.com/hifat/hifat-blog-api/internal/utils/response"
)

type authHandler struct {
	authService domain.AuthService
}

var validator utils.Validator

func NewAuthHandler(authService domain.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h authHandler) Register(ctx *gin.Context) {
	var req domain.RequestRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		response.FormErr(ctx, validator.Validate(err))
		return
	}

	res, err := h.authService.Register(req)
	if err != nil {
		response.InternalError(ctx, err)
		return
	}

	response.Created(ctx, response.SuccesResponse{
		Item: res,
	})
}
