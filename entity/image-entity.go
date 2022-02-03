package entity

import (
	null "gopkg.in/guregu/null.v4"
)

type Image struct {
	ImgId               null.Int
	EmpId               null.Int
	BrandId             null.Int
	CId                 null.Int
	UId                 null.Int
	RpAccId             null.Int
	ResId               null.Int
	ImgGuid             null.String
	FileName            null.String
	FilePathR           null.String
	FilePathM           null.String
	FilePathS           null.String
	MinDarkFileName     null.String
	MinDarkFilePath     null.String
	MaxDarkFileName     null.String
	MaxDarkFilePath     null.String
	MinLightFileName    null.String
	MinLightFilePath    null.String
	MaxLightFileName    null.String
	MaxLightFilePath    null.String
	CreatedDate         null.String
	ModifiedDate        null.String
	CreatedUId          null.Int
	ModifiedUId         null.Int
	SyncDateTime        null.String
	OptimisticLockField null.Int
	ResCatId            null.Int
	ProdId              null.Int
	TagId               null.Int
	GCRecord            null.Int
}
