package entity

import (
	null "gopkg.in/guregu/null.v4"
)

type Image struct {
	ImgId        null.Int
	EmpId        null.Int
	BrandId      null.Int
	ResId        null.Int
	CId          null.Int
	UId          null.Int
	RpAccId      null.Int
	ImgGuid      null.String
	FileName     null.String
	FilePath     null.String
	CreatedDate  null.String
	ModifiedDate null.String
	GCRecord     null.Int
}
