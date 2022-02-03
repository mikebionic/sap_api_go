package controller

import (
	"net/http"
	"sapgo/helper"
	"sapgo/service"

	"github.com/gin-gonic/gin"
)

type ResourceController interface {
	GetResource(context *gin.Context)
}
type resourceController struct {
	resourceService service.ResourceService
}

func NewResourceController(service service.ResourceService) ResourceController {
	return &resourceController{
		resourceService: service,
	}
}

func (c *resourceController) GetResource(context *gin.Context) {
	data := c.resourceService.GetResource()
	res := helper.BuildResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}
