package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "postgresql://404_owner:l9NPEM6TxtBs@ep-noisy-mountain-a1l9dhzu-pooler.ap-southeast-1.aws.neon.tech/404?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	return db, nil
}
