package entity

import "gopkg.in/guregu/null.v4"

type Category struct {
	ResCatId            null.Int
	ResCatVisibleIndex  null.Int
	IsMain              null.Bool
	ResCatName          null.String
	ResCatDesc          null.String
	ResCatIconName      null.String
	ResCatIconFilePath  null.String
	CreatedDate         null.String
	ModifiedDate        null.String
	CreatedUId          null.Int
	ModifiedUId         null.Int
	SyncDateTime        null.String
	OptimisticLockField null.Int
	GCRecord            null.Int
	ResOwnerCatId       null.Int
	ResCatGuid          null.String
}
