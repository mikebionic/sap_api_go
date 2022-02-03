package controller

import (
	"net/http"
	"sapgo/helper"
	"sapgo/service"

	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	GetCompany(context *gin.Context)
}
type companyController struct {
	companyService service.CompanyService
}

func NewCompanyController(service service.CompanyService) CompanyController {
	return &companyController{
		companyService: service,
	}
}

func (c *companyController) GetCompany(context *gin.Context) {
	data := c.companyService.GetCompany()
	res := helper.BuildResponse(true, "OK", data)
	context.JSON(http.StatusOK, res)
}
