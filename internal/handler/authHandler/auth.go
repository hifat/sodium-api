package authHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/handler/httpResponse"
	"github.com/hifat/hifat-blog-api/internal/utils/ernos"
	"github.com/hifat/hifat-blog-api/internal/utils/response"
	"github.com/hifat/hifat-blog-api/internal/utils/validity"
)

type authHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *authHandler {
	return &authHandler{authService}
}

// @Summary 	Register
// @Description Register
// @Tags 		Auth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} domain.ResponseRegister
// @Success 	422 {object} response.ErrorResponse "Unprocessable Entity"
// @Success 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/auth/register [post]
// @Param 		Body body domain.RequestRegister true "Register request"
func (h authHandler) Register(ctx *gin.Context) {
	var req domain.RequestRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	res, err := h.authService.Register(req)
	if err != nil {
		if e, ok := err.(ernos.Ernos); ok {
			if e.Code == ernos.C.DUPLICATE_RECORD {
				httpResponse.Conflict(ctx, err)
				return
			}
		}

		httpResponse.InternalError(ctx, err)
		return
	}

	httpResponse.Created(ctx, response.SuccesResponse{
		Item: res,
	})
}
