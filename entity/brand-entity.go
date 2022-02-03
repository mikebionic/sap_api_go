package entity

import "gopkg.in/guregu/null.v4"

type Brand struct {
	BrandId             null.Int
	BrandName           null.String
	BrandDesc           null.String
	BrandVisibleIndex   null.Int
	IsMain              null.Bool
	BrandLink1          null.String
	BrandLink2          null.String
	BrandLink3          null.String
	BrandLink4          null.String
	BrandLink5          null.String
	AddInf1             null.String
	AddInf2             null.String
	AddInf3             null.String
	AddInf4             null.String
	AddInf5             null.String
	AddInf6             null.String
	CreatedDate         null.String
	ModifiedDate        null.String
	CreatedUId          null.Int
	ModifiedUId         null.Int
	SyncDateTime        null.String
	OptimisticLockField null.Int
	GCRecord            null.Int
	BrandGuid           null.String
}
