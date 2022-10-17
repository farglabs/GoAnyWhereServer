package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"fileserver/db/data"
	"fileserver/db/schema"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	// Initialize the database file
	createDB()

	// Start the database with newly init file
	db, err := sql.Open("sqlite3", "db/init.db")
	data.CheckDBErr(err)

	// Create Schema
	schema.CreateTables(db)

	// PreSeed the data
	data.Seed(db)
}

func createDB() {
	dbFile, err := os.Create("db/init.db")
	if err != nil {
		fmt.Printf("Error %d", err)
		panic(err)
	}
	log.Println("Database File Created")
	dbFile.Close()
}
