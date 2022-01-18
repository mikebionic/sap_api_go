package controller

import (
	"admin/dto"
	"admin/helper"
	"admin/service"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type UploadController interface {
	UploadImage(c *gin.Context)
	GetFile(context *gin.Context)
}
type uploadController struct {
	uploadService service.UploadService
}

func NewUploadController(service service.UploadService) UploadController {
	return &uploadController{
		uploadService: service,
	}
}

func (c *uploadController) UploadImage(context *gin.Context) {
	var (
		errnum int
		num    int
		upload dto.Upload
	)
	errDTO := context.ShouldBind(&upload)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	form, err := context.MultipartForm()
	context.Request.MultipartReader()
	if err != nil {
		res := helper.BuildErrorResponse("get file error", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
	} else {
		files := form.File["files"]
		e := reflect.ValueOf(&upload).Elem()
		for i := 0; i < e.NumField(); i++ {
			if e.Field(i).Interface() != "" {
				guid := fmt.Sprintf("%s", e.Field(i).Interface())
				place := e.Type().Field(i).Name
				data := c.uploadService.Upload(files, guid, place)
				res := helper.BuildResponse(true, "Ok", data)
				for _, d := range data {
					num += 1
					if d.Error != "" {
						errnum += 1
					}
				}
				if num == errnum {
					context.JSON(http.StatusBadRequest, res)
				} else if num > errnum && errnum != 0 {
					context.JSON(http.StatusPartialContent, res)
				} else if errnum == 0 {
					context.JSON(http.StatusCreated, res)
				}
			}
		}
	}
}

func (c *uploadController) GetFile(context *gin.Context) {
	var get dto.GetFile
	err := context.BindJSON(&get)
	if err != nil {
		res := helper.BuildErrorResponse("Something Went Wrong", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
	}
	data, err := c.uploadService.GetFile(get)
	if err != nil {
		res := helper.BuildErrorResponse("Something Went Wrong", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
	} else {
		res := helper.BuildResponse(true, "Ok", data)
		context.JSON(http.StatusOK, res)
	}
}
