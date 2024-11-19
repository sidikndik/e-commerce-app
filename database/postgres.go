package database

import (
	"database/sql"
	"e-commerce-app/util"

	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(config util.Configuration) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s TimeZone=Asia/Jakarta",
		config.DBConfig.DBName, config.DBConfig.DBPassword, config.DBConfig.DBName, config.DBConfig.DBHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(50)

	return db, err
}
