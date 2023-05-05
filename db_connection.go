package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@host:hads1408@tcp(127.0.0.1:8080)/campusfora")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO data(id,input) VALUES(name,email)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the statement with some data
	res, err := stmt.Exec("Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Get the ID of the inserted row
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New user with ID %d inserted successfully", id)
}
