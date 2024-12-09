package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var JwtSecret []byte

// InitJwtSecret inisialisasi jwtSecret dari file .env
func InitJwtSecret() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading environment variables directly")
	}

	// Mengambil JWT_SECRET dari environment variables
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set in environment")
	}

	JwtSecret = []byte(secret)
}
