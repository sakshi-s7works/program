package database

import (
	"fmt"
	"log"
	"time"

	"github.com/sakshi-s7works/program/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	rdsEndpoint := "astro-db-test.c1ma9shuqqf3.us-east-2.rds.amazonaws.com"
	dbUser := "postgres"
	dbPassword := "C^494gD6H7blmv"
	dbName := "astro"
	// sslMode := "disable"
	dbPort := 5432
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", rdsEndpoint, dbUser, dbPassword, dbName, dbPort)

	// Configure Gorm with a connection pool
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	// Set up connection pool settings (adjust as needed)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error auto-migrating models: %v", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL")
}
