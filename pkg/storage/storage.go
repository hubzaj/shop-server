package storage

import (
	"database/sql"
	"fmt"
	"log"
)

type Storage struct {
	DB *sql.DB
}

func InitStorage(cfg *StorageConfig) *Storage {
	db, err := sql.Open("postgres", createConnectionStringWithConfig(cfg))
	if err != nil {
		log.Fatalf("error during connecting to postgres: %s", err)
	}
	return &Storage{
		DB: db,
	}
}

func createConnectionString(user, dbName, password, host, sslMode string, port int) string {
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		user,
		dbName,
		password,
		host,
		port,
		sslMode,
	)
}

func createConnectionStringWithConfig(cfg *StorageConfig) string {
	return createConnectionString(
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.Host,
		cfg.SSLMode,
		cfg.Port,
	)
}
