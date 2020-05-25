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

func GetAllGhsFromDb() {
	db := databaseConn()
	fetchAllGhsFromDB, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err)
	}
	var id int
	var title, chorus, stanza_1, stanza_2, stanza_3, stanza_4, stanza_5, stanza_6, stanza_7, stanza_8 string
	fetchAllGhsFromDB.Scan(&id, &title, &chorus, &stanza_1, &stanza_2, &stanza_3, &stanza_4, &stanza_5, &stanza_6, &stanza_7, &stanza_8)
}