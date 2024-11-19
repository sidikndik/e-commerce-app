package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	query := "user=postgres password=postgres dbname=postgres sslmode=disable host=localhost"
	db, err := sql.Open("postgres", query)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(50)

	return db, err
}
