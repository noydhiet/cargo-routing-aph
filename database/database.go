package database

import "database/sql"

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	//dbIP := "192.168.20.54"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
