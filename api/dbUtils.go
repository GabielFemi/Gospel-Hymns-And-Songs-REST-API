package ghs_api

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func databaseConn() *sql.DB {
	databaseDriver := "mysql"
	databaseUsername := "root"
	databaseName := "localhost"
	databasePassword := ""

	db, err := sql.Open(databaseDriver, databaseUsername+":"+databasePassword+"@/"+databaseName)

	if err != nil {
		panic(err)
	}
	return db
}
