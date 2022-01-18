package main

import (
	"admin/config"
	"admin/controller"
	"admin/middleware"
	"admin/repository"
	"admin/service"
	"admin/tools"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db               *sql.DB                     = config.SetupDatabaseConnection()
	uploadRepository repository.UploadRepository = repository.NewUploadRepository(db)
	uploadService    service.UploadService       = service.NewUploadService(uploadRepository)
	uploadController controller.UploadController = controller.NewUploadController(uploadService)
)

func main() {
	tools.EnvParser()
	r := gin.Default()
	r.Use(middleware.SetupCors())
	os.MkdirAll("./uploads/", os.ModePerm)

	r.POST("/", uploadController.UploadImage)
	r.GET("/", uploadController.GetFile)
	r.Run()
}
