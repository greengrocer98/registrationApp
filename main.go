package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbName = "registration.db"
var db *sql.DB

func main() {

	db, err := ConnectDB(dbName)
	if err != nil {
		log.Fatal("Couldn't connect to database")
	} else {
		log.Println("Connected to database")
	}

	defer db.Close()

}
