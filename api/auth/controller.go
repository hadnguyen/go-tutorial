package auth

import (
	"go-tutorial/api/auth/dto"
	"go-tutorial/arch/network"
	"go-tutorial/common"
	"go-tutorial/utils"

	"github.com/gin-gonic/gin"
)

type controller struct {
	network.BaseController
	common.ContextPayload
	service Service
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service Service,
) network.Controller {
	return &controller{
		BaseController: network.NewBaseController("/auth", authProvider, authorizeProvider),
		ContextPayload: common.NewContextPayload(),
		service:        service,
	}
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
	group.POST("/signup/basic", c.signUpBasicHandler)
	group.POST("/signin/basic", c.signInBasicHandler)
	group.POST("/token/refresh", c.tokenRefreshHandler)
	group.DELETE("/signout", c.Authentication(), c.signOutBasic)
}

func (c *controller) signUpBasicHandler(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, dto.EmptySignUpBasic())
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	data, err := c.service.SignUpBasic(body)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", data)
}

// @Summary		Sign in
// @Description	Sign in by email
// @Tags			auth
// @Accept			json
// @Produce		json
//
// @Security		BearerAuth
//
// @Param			X-API-KEY	header		string	true	"X-API-KEY is required"
// @Param			signIn			body		dto.SignInBasic	true	"sign in body"
// @Success		200				{object}	dto.UserAuth
// @Router			/auth/signin/basic [post]
func (c *controller) signInBasicHandler(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, dto.EmptySignInBasic())
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	dto, err := c.service.SignInBasic(body)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", dto)
}

func (c *controller) signOutBasic(ctx *gin.Context) {
	keystore := c.MustGetKeystore(ctx)

	err := c.service.SignOut(keystore)
	if err != nil {
		c.Send(ctx).InternalServerError("something went wrong", err)
		return
	}

	c.Send(ctx).SuccessMsgResponse("signout success")
}

func (c *controller) tokenRefreshHandler(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, dto.EmptyTokenRefresh())
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	authHeader := ctx.GetHeader(network.AuthorizationHeader)
	accessToken := utils.ExtractBearerToken(authHeader)

	dto, err := c.service.RenewToken(body, accessToken)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", dto)
}
