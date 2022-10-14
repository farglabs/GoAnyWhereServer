package db

import (
	"database/sql"
    "freshmanual.com/db/data"
	"freshmanual.com/db/schema"
	_ "github.com/mattn/go-sqlite3"
    "os"
    "log"
)

type initDB struct{}

func InitDB() (string, error){
    // Initialize the database file
    createDB()

    // Start the database with newly init file
	db, err := sql.Open("sqlite3", "db/init.db")
	data.CheckDBErr(err)
    
    // Create Schema
	schema.CreateTables(db)

    // PreSeed the data
    data.Seed(db)
    return "", &initDB{}
}

func (d *initDB) Error() string {
	return "Error Initializing Database"
}

func createDB() (string, error) {
	dbFile, err := os.Create("db/init.db")
	if err != nil {
		return "Something Went Wrong Creating the Database", err
	}
	log.Println("Database File Created")
	dbFile.Close()
	return "Database Successfully Created", &initDB{}
}