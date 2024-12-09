package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func InitDB() {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../.env")
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Failed to load environment variables")
	}

	// Menambahkan sslmode=disable
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connected successfully.")

	DB = database
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	sqlDB.Close()
}
