package entity

import (
	null "gopkg.in/guregu/null.v4"
)

type Company struct {
	CId                 null.Int
	CGuid               null.String
	CName               null.String
	CDesc               null.String
	CFullName           null.String
	AccInfId            null.Int
	CAddress            null.String
	CAddressLegal       null.String
	CLatitude           null.Float
	CLongitude          null.Float
	Phone1              null.String
	Phone2              null.String
	Phone3              null.String
	Phone4              null.String
	CPostalCode         null.String
	WebAddress          null.String
	CEmail              null.String
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
}
