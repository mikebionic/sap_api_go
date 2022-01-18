package service

import (
	"admin/dto"
	"admin/entity"
	"admin/repository"
	"admin/tools"
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

var imageSizes = [2]string{"S", "M"}
var jpegQualities = [2]int{10, 30}
var pngQualities = [2]string{"11", "9"}

type UploadService interface {
	Upload(data []*multipart.FileHeader, guid string, place string) []entity.Upload
	GetFile(d dto.GetFile) (entity.Image, error)
}
type uploadService struct {
	uploadRepository repository.UploadRepository
}

func NewUploadService(upload repository.UploadRepository) UploadService {
	return &uploadService{
		uploadRepository: upload,
	}
}
func (service *uploadService) Upload(data []*multipart.FileHeader, guid string, place string) []entity.Upload {
	var (
		upload  entity.Upload
		uploads []entity.Upload
	)
	tools.EnvParser()
	mb := os.Getenv("MAX_UPLOAD_SIZE")
	MAX_UPLOAD_SIZE, _ := strconv.ParseInt(mb, 0, 64)
	wg := sync.WaitGroup{}
	for _, fileHeader := range data {
		uuid := uuid.New().String()
		wg.Add(1)
		go func(fileHeader *multipart.FileHeader) {
			if fileHeader.Size > (MAX_UPLOAD_SIZE << 20) {
				upload.Error = fmt.Sprintf("File Size Too Big Should Be %dMB", MAX_UPLOAD_SIZE)
				upload.Name = fileHeader.Filename
				upload.TargetGuid = guid
				uploads = append(uploads, upload)
			} else {
				contentType, file := detectContentType(fileHeader)
				filetype := strings.SplitAfter(contentType, "/")
				switch contentType {
				case "image/png", "image/jpeg":
					os.MkdirAll(fmt.Sprintf("./uploads/%s/%s/images/S", place, uuid), os.ModePerm)
					os.MkdirAll(fmt.Sprintf("./uploads/%s/%s/images/M", place, uuid), os.ModePerm)
					os.MkdirAll(fmt.Sprintf("./uploads/%s/%s/images/R", place, uuid), os.ModePerm)
					err := realImageProcess(file, place, uuid, filetype[1])
					if err != nil {
						upload.Error = fmt.Sprintf("Get Image Error %s", err)
						upload.Name = fileHeader.Filename
						upload.TargetGuid = guid
						uploads = append(uploads, upload)
						os.RemoveAll(fmt.Sprintf("./uploads/%s/%s/images/", place, uuid))
					}
					err = imageProcess(place, contentType, uuid)
					if err != nil {
						upload.Error = fmt.Sprintf("Process Image Error %s", err)
						upload.Name = fileHeader.Filename
						upload.TargetGuid = guid
						uploads = append(uploads, upload)
						os.RemoveAll(fmt.Sprintf("./uploads/%s/%s/images/", place, uuid))
					} else {
						err := service.uploadRepository.Upload(upload, place)
						if err != nil {
							upload.Error = fmt.Sprintf("Database Error %s", err)
							upload.Name = fileHeader.Filename
							upload.TargetGuid = guid
							uploads = append(uploads, upload)
							os.RemoveAll(fmt.Sprintf("./uploads/%s/%s/images/", place, uuid))
						} else {
							upload.Error = ""
							upload.ImageGuid = uuid
							upload.Name = fileHeader.Filename
							upload.Path = fmt.Sprintf("./uploads/%s/%s/images/<FSIZE>/%s.%s", place, uuid, uuid, filetype[1])
							upload.TargetGuid = guid
							uploads = append(uploads, upload)
						}
					}
				default:
					upload.Error = "InvalidType"
					upload.Name = fileHeader.Filename
					upload.TargetGuid = guid
					uploads = append(uploads, upload)
				}
			}
			wg.Done()
		}(fileHeader)
	}
	wg.Wait()
	return uploads
}

func detectContentType(fileHeader *multipart.FileHeader) (string, multipart.File) {
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("FILE OPEN ERROR", err)
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println("READ ERROR", err)
	}
	filetype := http.DetectContentType(buff)
	file.Seek(0, io.SeekStart)
	return filetype, file
}

func imageProcess(place, contentType, uuid string) error {
	var w bytes.Buffer
	switch contentType {
	case "image/png":
		raw, err := os.Open(fmt.Sprintf("./uploads/%s/%s/images/R/%s.png", place, uuid, uuid))
		if err != nil {
			return err
		}
		defer raw.Close()
		img, err := png.Decode(raw)
		if err != nil {
			return err
		}
		png.Encode(&w, img)
		b := w.Bytes()
		for i := 0; i < 2; i++ {
			f, err := os.Create(fmt.Sprintf("./uploads/%s/%s/images/%s/%s.png", place, uuid, imageSizes[i], uuid))
			if err != nil {
				return err
			}
			compressed := compressPng(b, pngQualities[i])
			output, err := png.Decode(bytes.NewReader(compressed))
			if err != nil {
				return err
			}
			err = png.Encode(f, output)
			if err != nil {
				return err
			}
			f.Close()
		}
	case "image/jpeg":
		raw, err := os.Open(fmt.Sprintf("./uploads/%s/%s/images/R/%s.jpeg", place, uuid, uuid))
		if err != nil {
			return err
		}
		defer raw.Close()
		img, err := jpeg.Decode(raw)
		if err != nil {
			return err
		}
		for i := 0; i < 2; i++ {
			f, err := os.Create(fmt.Sprintf("./uploads/%s/%s/images/%s/%s.jpeg", place, uuid, imageSizes[i], uuid))
			if err != nil {
				return err
			}
			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpegQualities[i]})
			if err != nil {
				return err
			}
			f.Close()
		}
	default:
		return fmt.Errorf("invalid Type")
	}
	return nil
}
func compressPng(input []byte, val string) []byte {
	cmd := exec.Command("pngquant", "-", "--speed", val)
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	output := o.Bytes()
	return output
}
func realImageProcess(file multipart.File, place, uuid, filetype string) error {
	f, err := os.Create(fmt.Sprintf("./uploads/%s/%s/images/R/%s.%s", place, uuid, uuid, filetype))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return nil
}

func (service *uploadService) GetFile(d dto.GetFile) (entity.Image, error) {
	data, err := service.uploadRepository.GetFile(d)
	str := data.FilePath
	newStr := strings.Replace(str.String, "<FSize>", d.Size, 1)
	data.FilePath = null.NewString(newStr, true)
	return data, err
}
