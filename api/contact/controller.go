package contact

import (
	"go-tutorial/api/contact/dto"
	"go-tutorial/arch/network"
	"go-tutorial/utils"

	"github.com/gin-gonic/gin"
)

type controller struct {
	network.BaseController
	service Service
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service Service,
) network.Controller {
	return &controller{
		BaseController: network.NewBaseController("/contact", authProvider, authorizeProvider),
		service:        service,
	}
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
	group.POST("/", c.createMessageHandler)
}

func (c *controller) createMessageHandler(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, &dto.CreateMessage{})
	if err != nil {
		c.Send(ctx).BadRequestError(err.Error(), err)
		return
	}

	msg, err := c.service.SaveMessage(body)
	if err != nil {
		c.Send(ctx).InternalServerError("something went wrong", err)
		return
	}

	data, err := utils.MapTo[dto.InfoMessage](msg)
	if err != nil {
		c.Send(ctx).InternalServerError("something went wrong", err)
		return
	}

	c.Send(ctx).SuccessDataResponse("message received successfully!", data)
}
