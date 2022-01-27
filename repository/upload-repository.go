package repository

import (
	"database/sql"
	"fmt"
	"sapgo/dto"
	"sapgo/entity"
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
		row := db.connection.QueryRow(`SELECT "ImgId", "EmpId", "BrandId", "ResId", "CId", "UId",
		"RpAccId" , "ImgGuid", "FileName", "FilePath", "CreatedDate", "ModifiedDate", "GCRecord"
		FROM tbl_dk_image WHERE "ImgGuid"=$1;`, data.Guid)

		row.Scan(&image.ImgId, &image.EmpId, &image.BrandId, &image.ResId, &image.CId, &image.UId,
			&image.RpAccId, &image.ImgGuid, &image.FileName, &image.FilePath, &image.CreatedDate,
			&image.ModifiedDate, &image.GCRecord)
		if image.ImgGuid.String == "" {
			return entity.Image{}, fmt.Errorf("no such image")
		}
		return image, nil
	default:
		return entity.Image{}, fmt.Errorf("no such file")
	}
}
