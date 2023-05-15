package httpResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/response"
)

func Error(ctx *gin.Context, err any) {
	if e, ok := err.(ernos.Ernos); ok {
		ctx.AbortWithStatusJSON(e.Status, response.ErrorResponse{
			Error: e,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func BadRequest(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err.(error).Error())
}

func FormErr(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.HandleErr(err))
}

func Created(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusCreated, obj)
}

func Success(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}
