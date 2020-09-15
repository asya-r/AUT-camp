package models

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil
	}
	return db
}
