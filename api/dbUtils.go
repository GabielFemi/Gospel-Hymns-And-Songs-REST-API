package ghs_api

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func databaseConn() *sql.DB {
	databaseDriver := "mysql"
	databaseUsername := "root"
	databaseName := "ghs"
	databasePassword := ""

	db, err := sql.Open(databaseDriver, databaseUsername+":"+databasePassword+"@/"+databaseName)

	if err != nil {
		panic(err)
	}
	return db
}

func GetAllGhsFromDb() []GHS {
	db := databaseConn()
	fetchAllGhsFromDB, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err)
	}
	var id int
	var title, chorus, stanza1, stanza2, stanza3, stanza4, stanza5, stanza6, stanza7, stanza8 string
	var ghs GHS
	var ghsSlice []GHS
	for fetchAllGhsFromDB.Next() {
		err := fetchAllGhsFromDB.Scan(&id, &title, &chorus, &stanza1, &stanza2, &stanza3, &stanza4, &stanza5, &stanza6, &stanza7, &stanza8)
		if err != nil {
			panic(err)
		}
		ghs.Id = id
		ghs.Title = title
		ghs.Chorus = chorus
		ghs.Stanza1 = stanza1
		ghs.Stanza2 = stanza2
		ghs.Stanza3 = stanza3
		ghs.Stanza4 = stanza4
		ghs.Stanza5 = stanza5
		ghs.Stanza6 = stanza6
		ghs.Stanza8 = stanza8

		ghsSlice = append(ghsSlice, ghs)
	}
	return ghsSlice

}

func GetASingleGhs(ghsId interface{}) []GHS {
	db := databaseConn()
	fetchAllGhsFromDB, err := db.Query("SELECT * FROM songs WHERE id = ?", ghsId)
	if err != nil {
		panic(err)
	}
	var id int
	var title, chorus, stanza1, stanza2, stanza3, stanza4, stanza5, stanza6, stanza7, stanza8 string
	var ghs GHS
	var ghsSlice []GHS
	for fetchAllGhsFromDB.Next() {
		err := fetchAllGhsFromDB.Scan(&id, &title, &chorus, &stanza1, &stanza2, &stanza3, &stanza4, &stanza5, &stanza6, &stanza7, &stanza8)
		if err != nil {
			panic(err)
		}
		ghs.Id = id
		ghs.Title = title
		ghs.Chorus = chorus
		ghs.Stanza1 = stanza1
		ghs.Stanza2 = stanza2
		ghs.Stanza3 = stanza3
		ghs.Stanza4 = stanza4
		ghs.Stanza5 = stanza5
		ghs.Stanza6 = stanza6
		ghs.Stanza8 = stanza8

		ghsSlice = append(ghsSlice, ghs)
	}
	return ghsSlice
}