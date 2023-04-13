package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error any `json:"error"`
}

type ErrorMessageResponse struct {
	Message string `json:"message"`
}

type SuccesResponse struct {
	Item  any `json:"item,omitempty"`
	Items any `json:"items,omitempty"`
	Total any `json:"total,omitempty"`
}

func handleErr(err any) ErrorResponse {
	if _, ok := err.(error); ok {
		return ErrorResponse{
			Error: ErrorMessageResponse{
				Message: err.(error).Error(),
			},
		}
	}

	return ErrorResponse{
		Error: err,
	}
}

func FormErr(ctx *gin.Context, err any) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, handleErr(err))
}

func InternalError(ctx *gin.Context, err any) {
	log.Println(err.(error).Error())
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
		Error: ErrorMessageResponse{
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
