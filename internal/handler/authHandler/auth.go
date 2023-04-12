package authHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
)

type authHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h authHandler) Register(ctx *gin.Context) {
	var req domain.PayloadUser
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": utils.Validator(err),
		})
		return
	}

	res, err := h.authService.Register(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}
