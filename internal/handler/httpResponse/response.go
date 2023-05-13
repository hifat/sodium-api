package httpResponse

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/response"
)

func handleError(ctx *gin.Context, httpCode int, err error) {
	if _, ok := err.(ernos.Ernos); ok {
		log.Println(err.(ernos.Ernos).Error())
		ctx.AbortWithStatusJSON(httpCode, response.ErrorResponse{
			Error: ernos.Ernos{
				Message: err.(ernos.Ernos).Message,
				Code:    err.(ernos.Ernos).Code,
			},
		})
		return
	}

	log.Println(err.Error())
	ctx.AbortWithStatusJSON(httpCode, response.ErrorResponse{
		Error: ernos.Ernos{
			Message: err.Error(),
			Code:    "",
		},
	})
}

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

func InternalError(ctx *gin.Context, err any) {
	log.Println(err.(error).Error())
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
		Error: ernos.Ernos{
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

func Forbidden(ctx *gin.Context, err error) {
	handleError(ctx, http.StatusForbidden, err)
}

func Conflict(ctx *gin.Context, err error) {
	handleError(ctx, http.StatusConflict, err)
}

func Unauthorized(ctx *gin.Context, err error) {
	handleError(ctx, http.StatusUnauthorized, err)
}
