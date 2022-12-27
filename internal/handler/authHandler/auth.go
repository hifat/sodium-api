package authHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/domain"
)

type authHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h authHandler) Register(ctx *gin.Context) {
	var req domain.FormRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err, validateErors := h.authService.Register(req)
	if len(validateErors) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": validateErors,
		})
		return
	}

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
