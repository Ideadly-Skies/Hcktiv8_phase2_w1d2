package db

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// declare a global variable to hold the database connection
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:11111111@tcp(127.0.0.1:3306)/ftgo_phase2_w1d2")
	if err != nil {
		log.Print("Error connecting to the database: ", err)
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Print("Error pinging the database: ", err)
		log.Fatal(err)
	}
	
	fmt.Printf("Successfully connected to the database\n")

	return db, nil
}