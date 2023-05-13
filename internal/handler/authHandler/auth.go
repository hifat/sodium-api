package authHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/handler/httpResponse"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
	"github.com/hifat/sodium-api/internal/utils/response"
	"github.com/hifat/sodium-api/internal/utils/token"
	"github.com/hifat/sodium-api/internal/utils/validity"
)

type authHandler struct {
	authService authDomain.AuthService
}

func NewAuthHandler(authService authDomain.AuthService) *authHandler {
	return &authHandler{authService}
}

// @Summary		Register
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResponseRegister
// @Success		409 {object} response.ErrorResponse "Duplicate record"
// @Success		422 {object} response.ErrorResponse "Form validation error"
// @Success		500 {object} response.ErrorResponse "Internal server error"
// @Router		/auth/register [post]
// @Param		Body body authDomain.RequestRegister true "Register request"
func (h authHandler) Register(ctx *gin.Context) {
	fmt.Println(ctx.Request.UserAgent())
	fmt.Println(ctx.ClientIP())

	var req authDomain.RequestRegister
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	var res authDomain.ResponseRegister
	err = h.authService.Register(req, &res)
	if err != nil {
		httpResponse.Error(ctx, err)
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
// @Success		200 {object} authDomain.ResponseRefreshToken
// @Success		401 {object} response.ErrorResponse "Username or password is incorect"
// @Success		422 {object} response.ErrorResponse "Form validation error"
// @Success		500 {object} response.ErrorResponse "Internal server error"
// @Router		/auth/login [post]
// @Param		Body body authDomain.RequestLogin true "Register request"
func (h authHandler) Login(ctx *gin.Context) {
	var req authDomain.RequestLogin
	err := ctx.ShouldBind(&req)
	if err != nil {
		httpResponse.FormErr(ctx, validity.Validate(err))
		return
	}

	req.Agent = ctx.Request.UserAgent()
	req.ClientIP = utype.IP(ctx.ClientIP())

	res := authDomain.ResponseRefreshToken{}
	err = h.authService.Login(req, &res)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	httpResponse.Success(ctx, response.SuccesResponse{
		Item: res,
	})
}

// @Summary		Get Refresh Token
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Success		200 {object} authDomain.ResponseRefreshToken
// @Success		401 {object} response.ErrorResponse "Username or password is incorect"
// @Success		422 {object} response.ErrorResponse "Form validation error"
// @Success		500 {object} response.ErrorResponse "Internal server error"
// @Router		/auth/token/refresh [post]
// @Param		Body body authDomain.RequestToken true "Register request"
func (h authHandler) CreateRefreshToken(ctx *gin.Context) {
	credentials := ctx.MustGet("credentials").(*token.Payload)
	req := authDomain.RequestCreateRefreshToken{
		UserID:   credentials.UserID,
		Agent:    ctx.Request.UserAgent(),
		ClientIP: utype.IP(ctx.ClientIP()),
	}

	res, err := h.authService.CreateRefreshToken(req)
	if err != nil {
		httpResponse.Error(ctx, err)
		return
	}

	httpResponse.Success(ctx, response.SuccesResponse{
		Item: res,
	})
}
