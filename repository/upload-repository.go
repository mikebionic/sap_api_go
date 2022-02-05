package repository

import (
	"database/sql"
	"fmt"
	"sapgo/dto"
	"sapgo/entity"

	"gopkg.in/guregu/null.v4"
)

//INSERT INTO tbl_dk_image ("ImgGuid","RpAccId") VALUES ( '69d885d8-022e-484b-9cf2-e413ca27294c', (SELECT "RpAccId" FROM tbl_dk_rp_acc WHERE "RpAccGuid"='954edbe5-049a-455d-a7ab-0729a18affbf' LIMIT 1));
type UploadRepository interface {
	Upload(data entity.Upload, place string) error
	GetFile(data dto.GetFile) (entity.Image, error)
}
type uploadConnection struct {
	connection *sql.DB
}

func NewUploadRepository(db *sql.DB) UploadRepository {
	return &uploadConnection{
		connection: db,
	}
}

func (db *uploadConnection) Upload(data entity.Upload, place string) error {
	switch place {
	case "RpAccGuid":
		fmt.Println(data.TargetGuid)
		_, err := db.connection.Exec(`INSERT INTO tbl_dk_image ("ImgGuid","RpAccId","FilePath", "FileName") 
		VALUES ($1, (SELECT "RpAccId" FROM tbl_dk_rp_acc WHERE "RpAccGuid"=$2 LIMIT 1),$3,$4);`,
			data.ImageGuid, data.TargetGuid, data.Path, data.Name)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("column Error")
	}
	return nil
}

func (db *uploadConnection) GetFile(data dto.GetFile) (entity.Image, error) {
	switch data.File {
	case "image":
		var image entity.Image
		row := db.connection.QueryRow(`SELECT "ImgId", "EmpId", "BrandId", "CId", "UId",
		"RpAccId", "ResId", "ImgGuid", "FileName", "FilePath", "MinDarkFileName", "MinDarkFilePath",
		"MaxDarkFileName", "MaxDarkFilePath", "MinLightFileName", "MinLightFilePath", "MaxLightFileName",
		"MaxLightFilePath", "CreatedDate", "ModifiedDate", "CreatedUId", "ModifiedUId", "SyncDateTime", "OptimisticLockField",
		"GCRecord", "ResCatId", "ProdId", "TagId" FROM tbl_dk_image WHERE "ImgGuid"=$1;`, data.Guid)

		row.Scan(&image.ImgId, &image.EmpId, &image.BrandId, &image.CId, &image.UId, &image.RpAccId,
			&image.ResId, &image.ImgGuid, &image.FileName, &image.FilePathR, &image.MinDarkFileName, &image.MinDarkFilePath,
			&image.MaxDarkFileName, &image.MaxDarkFilePath, &image.MinLightFileName, &image.MinLightFilePath,
			&image.MaxLightFileName, &image.MaxLightFilePath, &image.CreatedDate, &image.ModifiedDate,
			&image.CreatedUId, &image.ModifiedUId, &image.SyncDateTime, &image.OptimisticLockField, &image.GCRecord,
			&image.ResCatId, &image.ProdId, &image.TagId)
		if image.ImgGuid.String == "" {
			return entity.Image{}, fmt.Errorf("no such image")
		}
		sizedImage := ImageToSizes(image)
		return sizedImage, nil
	default:
		return entity.Image{}, fmt.Errorf("no such file")
	}
}

func ImageToSizes(d entity.Image) entity.Image {
	d.FilePathM = null.NewString(("http://127.0.0.1:8080/get-image/M/" + d.ImgGuid.String), true)
	d.FilePathR = null.NewString(("http://127.0.0.1:8080/get-image/R/" + d.ImgGuid.String), true)
	d.FilePathS = null.NewString(("http://127.0.0.1:8080/get-image/S/" + d.ImgGuid.String), true)
	return d
}
