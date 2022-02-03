package repository

import (
	"database/sql"
	"fmt"
	"sapgo/entity"
)

type CompanyRepository interface {
	GetCompany() []entity.Company
}
type companyConnection struct {
	connection *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyConnection{
		connection: db,
	}
}

func (db *companyConnection) GetCompany() []entity.Company {
	var (
		companies []entity.Company
		company   entity.Company
	)
	rows, err := db.connection.Query(`SELECT "CId", "CGuid", "CName", "CDesc", "CFullName", "AccInfId",
	"CAddress", "CAddressLegal", "CLatitude", "CLongitude", "Phone1", "Phone2", "Phone3", "Phone4",
	"CPostalCode", "WebAddress", "CEmail", "AddInf1", "AddInf2", "AddInf3", "AddInf4", "AddInf5", 
	"AddInf6", "CreatedDate", "ModifiedDate", "CreatedUId", "ModifiedUId", "SyncDateTime", "OptimisticLockField",
	"GCRecord" FROM tbl_dk_company;`)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&company.CId, &company.CGuid, &company.CName, &company.CDesc, &company.CFullName,
			&company.AccInfId, &company.CAddress, &company.CAddressLegal, &company.CLatitude, &company.CLongitude,
			&company.Phone1, &company.Phone2, &company.Phone3, &company.Phone4, &company.CPostalCode,
			&company.WebAddress, &company.CEmail, &company.AddInf1, &company.AddInf2, &company.AddInf3,
			&company.AddInf4, &company.AddInf5, &company.AddInf6, &company.CreatedDate, &company.ModifiedDate,
			&company.CreatedUId, &company.ModifiedUId, &company.SyncDateTime, &company.OptimisticLockField,
			&company.GCRecord)
		if err != nil {
			fmt.Println(err)
		}
		companies = append(companies, company)
	}
	defer rows.Close()
	return companies
}
