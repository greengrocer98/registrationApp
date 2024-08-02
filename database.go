package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

func ConnectDB(dbName string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	_, err = os.Stat(dbName)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Creating new database...")
		db, err = CreateDB(dbName) // create new database
	} else {
		log.Println("Opening database...")
		db, err = sql.Open("sqlite3", dbName) // open existing database
	}
	return db, err
}

func CreateDB(dbName string) (*sql.DB, error) {
	var err error
	var db *sql.DB

	createUserTable := `CREATE TABLE User (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"email" TEXT,		
		"password" TEXT,		
		"first_name" TEXT,
		"last_name" TEXT,
		"sur_name" TEXT,		
		"phone" INTEGER,		
		"role" TEXT,		
		"rating" INTEGER		
	  );`

	createTournamentTable := `CREATE TABLE Tournament (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,		
		"time" TEXT,		
		"capacity" INTEGER,
		"users_id" TEXT
	  );`

	createUserTourShipTable := `CREATE TABLE UserTourShip (
		tour_id INTEGER NOT NULL REFERENCES Tournament(id),
		user_id INTEGER NOT NULL REFERENCES User(id),
    	PRIMARY KEY(tour_id, user_id)
	  );`

	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return db, err
	}

	log.Println("Creating User table...")
	statement, err := db.Prepare(createUserTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("User table created")

	log.Println("Creating Tournament table...")
	statement, err = db.Prepare(createTournamentTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Tournament table created")

	log.Println("Creating UserTourShip table...")
	statement, err = db.Prepare(createUserTourShipTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("UserTourShip table created")
	return db, err
}
