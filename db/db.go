package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Retrieve database connection details from environment variables
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("3306")

	// Construct the DSN (Data Source Name) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	var db *gorm.DB
	retries := 5 // Number of retries for database connection

	for retries > 0 {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: true,                                // Enable prepared statements for improved performance
			Logger:      logger.Default.LogMode(logger.Info), // Enable detailed logs for debugging
		})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database, retries left: %d", retries)
		time.Sleep(2 * time.Second) // Wait before retrying
		retries--
	}

	if err != nil {
		log.Panic("Could not connect to the database: " + err.Error())
	}

	// Retrieve the underlying sql.DB instance for connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("Failed to retrieve database instance: " + err.Error())
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(50)                  // Allow up to 50 idle connections
	sqlDB.SetMaxOpenConns(1000)                // Allow up to 1000 open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Limit each connection's lifetime to 30 minutes

	// Health check for the connection
	if err := sqlDB.Ping(); err != nil {
		log.Panic("Failed to ping the database: " + err.Error())
	}

	log.Println("Database connection successfully established.")
	return db
}
