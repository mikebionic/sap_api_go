package controller

import (
	"net/http"
	"sapgo/helper"
	"sapgo/service"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetCategory(context *gin.Context)
}
type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: service,
	}
}

func (c *categoryController) GetCategory(context *gin.Context) {
	data := c.categoryService.GetCategory()
	res := helper.BuildResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}
