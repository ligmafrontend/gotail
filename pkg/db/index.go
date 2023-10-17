package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

func Init() {
	// open local sqlite database db.sqlite
	var err error
	DB, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// create table if it doesn't exist
	initSQL, err := os.ReadFile("pkg/db/init.sql")
	if err != nil {
		log.Fatal("Error reading init.sql: ", err)
	}

	statements := strings.Split(string(initSQL), ";")

	// Execute each statement
	for _, statement := range statements {
		if strings.TrimSpace(statement) != "" {
			_, err := DB.Exec(statement)
			if err != nil {
				log.Fatal("Error executing SQL statement: ", err)
			}
		}
	}
}
