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
	"strings"

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
	r.Static("/get-img", "./uploads")
	r.Use(middleware.SetupCors())

	os.MkdirAll("./uploads/", os.ModePerm)

	r.GET("/get-image/:size/:file", func(c *gin.Context) {
		size := c.Param("size")
		file := c.Param("file")
		guid := strings.Split(file, ".")
		str := resourceRepository.GetImage(guid[0])
		newPath := strings.Replace(str.String, "<FSize>", size, 1)
		c.File(newPath)
	})
	r.GET("/", uploadController.GetFile)
	r.POST("/", uploadController.UploadImage)

	r.GET("/company", companyController.GetCompany)

	r.GET("/resource", resourceController.GetResource)
	r.Run()
}
