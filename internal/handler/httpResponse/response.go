package httpResponse

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/response"
)

func FormErr(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.HandleErr(err))
}

func InternalError(ctx *gin.Context, err any) {
	log.Println(err.(error).Error())
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
		Error: response.ErrorMessageResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		},
	})
}

func Created(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusCreated, obj)
}

func Success(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, obj)
}

func Conflict(ctx *gin.Context, err any) {
	log.Println(err.(ernos.Ernos).Error())
	ctx.AbortWithStatusJSON(http.StatusConflict, response.ErrorResponse{
		Error: response.ErrorMessageResponse{
			Message: err.(ernos.Ernos).Message,
			Code:    err.(ernos.Ernos).Code,
		},
	})
}
