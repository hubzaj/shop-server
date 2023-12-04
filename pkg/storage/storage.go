package storage

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func InitStorage() {
	var err error
	Database, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/mydb?sslmode=disable")
	if err != nil {
		log.Fatalf("error during connecting to postgres: %s", err)
	}
}
