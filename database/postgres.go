package database

import (
	"e-commerce-app/util"
	"log"
	"os"

	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config util.Configuration) (*gorm.DB, error) {
	// Format connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s TimeZone=Asia/Jakarta",
		config.DBConfig.DBName, config.DBConfig.DBPassword, config.DBConfig.DBName, config.DBConfig.DBHost)

	// Setup logger for GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	// Open a connection to the PostgreSQL databas
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// Convert to *sql.DB for setting connection options
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool options
	sqlDB.SetConnMaxIdleTime(time.Duration(config.DBConfig.DBMaxIdleTime) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(config.DBConfig.DBMaxLifeTime) * time.Hour)
	sqlDB.SetMaxIdleConns(config.DBConfig.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DBConfig.DBMaxOpenConns)

	return db, err
}
