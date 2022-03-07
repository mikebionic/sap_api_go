package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sapgo/entity"
	"sapgo/tools"
)

type CategoryRepository interface {
	GetCategory() []entity.Category
}
type categoryConneciton struct {
	connection *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryConneciton{
		connection: db,
	}
}

func (db *categoryConneciton) GetCategory() []entity.Category {
	tools.EnvParser()
	var (
		category     entity.Category
		categories   []entity.Category
		visibleIndex string
	)
	if visibleIndex = os.Getenv("CATEGORY_VISIBLE_INDEX"); visibleIndex == "" {
		visibleIndex = "-1"
	}
	rows, err := db.connection.Query(`SELECT "ResCatId", "ResCatVisibleIndex","IsMain","ResCatName","ResCatDesc",
		"ResCatIconName", "ResCatIconFilePath", "CreatedDate", "ModifiedDate", "CreatedUId", "ModifiedUId", "SyncDateTime",
		"OptimisticLockField", "GCRecord", "ResOwnerCatId", "ResCatGuid" FROM tbl_dk_res_category WHERE "ResCatVisibleIndex" > $1;`, visibleIndex)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&category.ResCatId, &category.ResCatVisibleIndex, &category.IsMain,
			&category.ResCatName, &category.ResCatDesc, &category.ResCatIconName, &category.ResCatIconFilePath,
			&category.CreatedDate, &category.ModifiedDate, &category.CreatedUId, &category.ModifiedUId,
			&category.SyncDateTime, &category.OptimisticLockField, &category.GCRecord, &category.ResOwnerCatId,
			&category.ResCatGuid)
		if err != nil {
			fmt.Println(err)
		}
		categories = append(categories, category)
	}
	return categories
}
