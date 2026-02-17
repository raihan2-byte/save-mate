package database

import (
	"database/sql"
	"log"
)

func InitDatabase() *sql.DB {

	db, err := ConnectionDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	RunMigration(db)

	return db
}
