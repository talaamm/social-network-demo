package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// GetDB returns the database connection
func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		log.Println(err)
	}
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
