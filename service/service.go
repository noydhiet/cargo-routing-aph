package service

import (
	"database/sql"
	"fmt"
	"kit/log"
	"os"
	"strings"

	dt "cargo-routing/datastruct"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "db_go"
	dbIP := "localhost"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func HelloWorld(name string) string {

	var helloOutput string

	helloOutput = fmt.Sprintf("Hello, $s ", name)

	return helloOutput

}

func HelloDaerah(name string, kelamin string, asal string) string {

	var helloOutput string

	switch strings.ToUpper(asal) {
	case "JAKARTA":
		{
			if strings.ToUpper(kelamin) == "PRIA" {
				helloOutput = fmt.Sprintf("Hi, MR ", name)
			} else {
				helloOutput = fmt.Sprintf("Hi, MRS  ", name)
			}
		}
	case "BANDUNG":
		{
			if strings.ToUpper(kelamin) == "PRIA" {
				helloOutput = fmt.Sprintf("Wilujeng, MR ", name)
			} else {
				helloOutput = fmt.Sprintf("Wilujeng, MRS  ", name)
			}
		}
	case "MEDAN":
		{
			if strings.ToUpper(kelamin) == "PRIA" {
				helloOutput = fmt.Sprintf("Horas, MR ", name)
			} else {
				helloOutput = fmt.Sprintf("Horas, MRS  ", name)
			}
		}

	}

	//helloOutput = fmt.Sprintf("Hello, $s ", name)

	return helloOutput

}

func GetStatusDelivery(del dt.Delivery) []dt.Delivery {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Checking Database")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM t_mtr_delivery where route_id = ?", del.ID_ROUTE)
	if err != nil {
		panic(err.Error())
	}
	delv := dt.Delivery{}
	res := []dt.Delivery{}

	for selDB.Next() {
		var idDelivery, routeId, itenaryId int
		var routingStatus, transportStatus, lastKnownLoc string
		err = selDB.Scan(&idDelivery, &routingStatus, &transportStatus, &lastKnownLoc, &itenaryId, &routeId)
		if err != nil {
			panic(err.Error())
		}
		delv.ID_DELIVERY = idDelivery
		delv.ROUTING_STATUS = routingStatus
		delv.TRANSPORT_STATUS = transportStatus
		delv.LAST_KNOWN_LOCATION = lastKnownLoc
		delv.ID_ITENARY = itenaryId
		delv.ID_ROUTE = routeId
		res = append(res, delv)
	}

	return res

}
