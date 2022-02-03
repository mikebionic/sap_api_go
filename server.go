package main

import (
	"database/sql"
	"os"
	"sapgo/config"
	"sapgo/controller"
	"sapgo/middleware"
	"sapgo/repository"
	"sapgo/service"
	"sapgo/tools"

	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB = config.SetupDatabaseConnection()

	uploadRepository repository.UploadRepository = repository.NewUploadRepository(db)
	uploadService    service.UploadService       = service.NewUploadService(uploadRepository)
	uploadController controller.UploadController = controller.NewUploadController(uploadService)

	companyRepository repository.CompanyRepository = repository.NewCompanyRepository(db)
	companyService    service.CompanyService       = service.NewCompanyService(companyRepository)
	companyController controller.CompanyController = controller.NewCompanyController(companyService)

	resourceRepository repository.ResourceRepository = repository.NewResourceRepository(db)
	resourceService    service.ResourceService       = service.NewResourceService(resourceRepository)
	resourceController controller.ResourceController = controller.NewResourceController(resourceService)
)

func main() {
	tools.EnvParser()

	r := gin.Default()
	r.Use(middleware.SetupCors())

	os.MkdirAll("./uploads/", os.ModePerm)

	r.GET("/", uploadController.GetFile)
	r.POST("/", uploadController.UploadImage)

	r.GET("/company", companyController.GetCompany)

	r.GET("/resource", resourceController.GetResource)
	r.Run()
}
