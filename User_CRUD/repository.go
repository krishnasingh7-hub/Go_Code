package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	// Open database connection
	var err error
	db, err = sql.Open("mysql", "root:Joie_Vibre@0!9@tcp(localhost:3306)/user_crud")
	if err != nil {
		return err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to the database")

	return nil
}
