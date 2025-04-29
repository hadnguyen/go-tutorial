package user

import (
	coredto "go-tutorial/arch/dto"
	"go-tutorial/arch/network"
	"go-tutorial/common"

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
		BaseController: network.NewBaseController("/profile", authProvider, authorizeProvider),
		ContextPayload: common.NewContextPayload(),
		service:        service,
	}
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
	group.GET("/id/:id", c.getPublicProfileHandler)
	private := group.Use(c.Authentication())
	private.GET("/mine", c.getPrivateProfileHandler)
}

func (c *controller) getPublicProfileHandler(ctx *gin.Context) {
	mongoId, err := network.ReqParams(ctx, coredto.EmptyMongoId())
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	data, err := c.service.GetUserPublicProfile(mongoId.ID)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", data)
}

// @Summary		Get profile
// @Description	Get private profile
// @Tags			auth
// @Accept			json
// @Produce		json
//
// @Security		BearerAuth
//
// @Param			X-API-KEY	header		string	true	"X-API-KEY is required"
// @Success		200			{object}	dto.InfoPrivateUser
// @Router			/profile/mine [get]
func (c *controller) getPrivateProfileHandler(ctx *gin.Context) {
	user := c.MustGetUser(ctx)

	data, err := c.service.GetUserPrivateProfile(user)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

	c.Send(ctx).SuccessDataResponse("success", data)
}
