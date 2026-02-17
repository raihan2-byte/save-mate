package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/savemate?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
