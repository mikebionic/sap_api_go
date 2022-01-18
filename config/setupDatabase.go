package config

import (
	"admin/tools"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

var (
	dbType = ""
)

//Get String DSN string for database
func connSting() string {
	tools.EnvParser()
	dbType = os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbParam := os.Getenv("DB_PARAMS")
	//connectString := "sqlserver://SBM:myP%4055w0rd@VM17:1433?database=AE&connection+timeout=30"
	switch dbType {
	case "sqlserver":
		return fmt.Sprintf("%s://%s:%s@%s?database=%s&%s", dbType, dbUser, dbPass, dbHost, dbName, dbParam)
	default:
		return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", dbType, dbUser, dbPass, dbHost, dbName)
	}
}

//Setup Connection With database
func SetupDatabaseConnection() *sql.DB {
	db, err := sql.Open(dbType, connSting())
	if err != nil {
		fmt.Print(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("connected")
	return db
}
