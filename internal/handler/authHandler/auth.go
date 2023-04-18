package authHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/domain"
	"github.com/hifat/sodium-api/internal/handler/httpResponse"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
	"github.com/hifat/sodium-api/internal/utils/response"
	"github.com/hifat/sodium-api/internal/utils/validity"
)

type authHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *authHandler {
	return &authHandler{authService}
}

// @Summary		Register
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} domain.ResponseRegister
// @Success		409 {object} response.ErrorResponse "Duplicate record"
// @Success		422 {object} response.ErrorResponse "Form validation error"
// @Success		500 {object} response.ErrorResponse "Internal server error"
// @Router		/auth/register [post]
// @Param		Body body domain.RequestRegister true "Register request"
func (h authHandler) Register(ctx *gin.Context) {
	fmt.Println(ctx.Request.UserAgent())
	fmt.Println(ctx.ClientIP())

	var req domain.RequestRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	var res domain.ResponseRegister
	err = h.authService.Register(req, &res)
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

// @Summary		Login
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} domain.RequestLogin
// @Success		401 {object} response.ErrorResponse "Username or password is incorect"
// @Success		422 {object} response.ErrorResponse "Form validation error"
// @Success		500 {object} response.ErrorResponse "Internal server error"
// @Router		/auth/login [post]
// @Param		Body body domain.ResponseLogin true "Register request"
func (h authHandler) Login(ctx *gin.Context) {
	var req domain.RequestLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	req.Agent = ctx.Request.UserAgent()
	req.ClientIP = utype.IP(ctx.ClientIP())

	res := domain.ResponseLogin{}
	err = h.authService.Login(req, &res)
	if err != nil {
		if e, ok := err.(ernos.Ernos); ok {
			if e.Code == ernos.C.INVALID_CREDENTIALS {
				httpResponse.Unauthorized(ctx, err)
				return
			}
		}

		httpResponse.InternalError(ctx, err)
		return
	}

	httpResponse.Success(ctx, response.SuccesResponse{
		Item: res,
	})
}
