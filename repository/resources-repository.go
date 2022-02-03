package repository

import (
	"database/sql"
	"fmt"
	"sapgo/entity"
)

type ResourceRepository interface {
	GetResource() []entity.Resource
}
type resourceConnection struct {
	connection *sql.DB
}

func NewResourceRepository(db *sql.DB) ResourceRepository {
	return &resourceConnection{
		connection: db,
	}
}

func (db *resourceConnection) GetResource() []entity.Resource {
	var (
		resources []entity.Resource
		resource  entity.Resource
		images    []entity.Image
		image     entity.Image
		brand     entity.Brand
		category  entity.Category
	)
	rows, err := db.connection.Query(
		`SELECT res."ResId", res."CId", res."DivId", res."ResCatId", res."UnitId", res."BrandId",
		res."UsageStatusId", res."ResTypeId", res."ResMainImgId", res."ResMakerId", res."ResLastVendorId", 
		res."ResGuid", res."ResRegNo", res."ResName", res."ResDesc", res."ResFullDesc", res."ResWidth", 
		res."ResHeight", res."ResLength", res."ResWeight", res."ResProductionOnSale", res."ResMinSaleAmount", 
		res."ResMaxSaleAmount", res."ResMinSalePrice", res."ResMaxSalePrice", res."AddInf1", res."AddInf2", 
		res."AddInf3", res."AddInf4", res."AddInf5", res."AddInf6", res."CreatedDate", res."ModifiedDate", 
		res."CreatedUId", res."ModifiedUId", res."SyncDateTime", res."OptimisticLockField", res."GCRecord", res."TagId",
		res."IsMain", res."ResVisibleIndex", res."ResViewCnt", br."BarcodeVal", b."BrandId", b."BrandName",
		b."BrandDesc", b."BrandVisibleIndex", b."IsMain", b."BrandLink1", b."BrandLink2", b."BrandLink3", 
		b."BrandLink4", b."BrandLink5", b."AddInf1", b."AddInf2", b."AddInf3", b."AddInf4", b."AddInf5",
		b."AddInf6", b."CreatedDate", b."ModifiedDate", b."CreatedUId", b."ModifiedUId", b."SyncDateTime", b."OptimisticLockField",
		b."GCRecord", b."BrandGuid", rc."ResCatId", rc."ResCatVisibleIndex", rc."IsMain", rc."ResCatName", 
		rc."ResCatIconName", rc."ResCatIconFilePath", rc."CreatedDate", rc."ModifiedDate", 
		rc."CreatedUId", rc."ModifiedUId", rc."SyncDateTime", rc."OptimisticLockField",
		rc."GCRecord", rc."ResOwnerCatId", rc."ResCatGuid", p."ResPriceValue"
		FROM tbl_dk_resource as res 
		INNER JOIN  (SELECT * FROM tbl_dk_res_price WHERE "ResPriceTypeId"='2') as p 
		ON res."ResId"=p."ResId" LEFT OUTER JOIN tbl_dk_brand as b 
		ON res."BrandId"=b."BrandId" LEFT JOIN tbl_dk_barcode as br 
		ON res."ResId"=br."ResId" LEFT OUTER JOIN tbl_dk_res_category as rc 
		ON res."ResCatId"=rc."ResCatId";`)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&resource.ResId, &resource.CId, &resource.DivId, &resource.ResCatId, &resource.UnitId, &resource.BrandId,
			&resource.UsageStatusId, &resource.ResTypeId, &resource.ResMainImgId, &resource.ResMakerId, &resource.ResLastVendorId,
			&resource.ResGuid, &resource.ResRegNo, &resource.ResName, &resource.ResDesc, &resource.ResFullDesc, &resource.ResWidth,
			&resource.ResHeight, &resource.ResLength, &resource.ResWeight, &resource.ResProductionOnSale, &resource.ResMinSaleAmount,
			&resource.ResMaxSaleAmount, &resource.ResMinSalePrice, &resource.ResMaxSalePrice, &resource.AddInf1, &resource.AddInf2,
			&resource.AddInf3, &resource.AddInf4, &resource.AddInf5, &resource.AddInf6, &resource.CreatedDate, &resource.ModifiedDate,
			&resource.CreatedUId, &resource.ModifiedUId, &resource.SyncDateTime, &resource.OptimisticLockField, &resource.GCRecord, &resource.TagId,
			&resource.IsMain, &resource.ResVisibleIndex, &resource.ResViewCnt, &resource.BarcodeValue, &brand.BrandId, &brand.BrandName,
			&brand.BrandDesc, &brand.BrandVisibleIndex, &brand.IsMain, &brand.BrandLink1, &brand.BrandLink2, &brand.BrandLink3,
			&brand.BrandLink4, &brand.BrandLink5, &brand.AddInf1, &brand.AddInf2, &brand.AddInf3, &brand.AddInf4, &brand.AddInf5,
			&brand.AddInf6, &brand.CreatedDate, &brand.ModifiedDate, &brand.CreatedUId, &brand.ModifiedUId, &brand.SyncDateTime, &brand.OptimisticLockField,
			&brand.GCRecord, &brand.BrandGuid, &category.ResCatId, &category.ResCatVisibleIndex, &category.IsMain, &category.ResCatName,
			&category.ResCatIconName, &category.ResCatIconFilePath, &category.CreatedDate, &category.ModifiedDate,
			&category.CreatedUId, &category.ModifiedUId, &category.SyncDateTime, &category.OptimisticLockField,
			&category.GCRecord, &category.ResOwnerCatId, &category.ResCatGuid, &resource.Price)
		if err != nil {
			fmt.Println(err)
		}
		resource.Brand = brand
		resource.Category = category
		irows, err := db.connection.Query(`SELECT "ImgId", "EmpId", "BrandId", "CId", "UId",
		"RpAccId", "ResId", "ImgGuid", "FileName", "FilePath", "MinDarkFileName", "MinDarkFilePath",
		"MaxDarkFileName", "MaxDarkFilePath", "MinLightFileName", "MinLightFilePath", "MaxLightFileName",
		"MaxLightFilePath", "CreatedDate", "ModifiedDate", "CreatedUId", "ModifiedUId", "SyncDateTime", "OptimisticLockField",
		"GCRecord", "ResCatId", "ProdId", "TagId" FROM tbl_dk_image WHERE "ResId"=$1;`, resource.ResId)
		if err != nil {
			fmt.Println(err)
		}
		for irows.Next() {
			err = irows.Scan(&image.ImgId, &image.EmpId, &image.BrandId, &image.CId, &image.UId, &image.RpAccId,
				&image.ResId, &image.ImgGuid, &image.FileName, &image.FilePathR, &image.MinDarkFileName, &image.MinDarkFilePath,
				&image.MaxDarkFileName, &image.MaxDarkFilePath, &image.MinLightFileName, &image.MinLightFilePath,
				&image.MaxLightFileName, &image.MaxLightFilePath, &image.CreatedDate, &image.ModifiedDate,
				&image.CreatedUId, &image.ModifiedUId, &image.SyncDateTime, &image.OptimisticLockField, &image.GCRecord,
				&image.ResCatId, &image.ProdId, &image.TagId)
			if err != nil {
				fmt.Println(err)
			}
			sizedImage := ImageToSizes(image)
			images = append(images, sizedImage)
		}
		resource.Image = images
		resources = append(resources, resource)
	}
	return resources
}
