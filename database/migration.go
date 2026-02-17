package database

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {

	for _, script := range migrationScripts {

		_, err := db.Exec(script)
		if err != nil {
			log.Fatal(err)
		}

	}
}
