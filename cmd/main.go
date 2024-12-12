package main

import (
	"github.com/joho/godotenv"
	"live-code-2-XioweL/config"
	"live-code-2-XioweL/internal/routes"
	"log"
	"os"
	"path/filepath"
)

var jwtSecret []byte

func init() {
	// Muat .env
	//if err := godotenv.Load(); err != nil {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, "../.env")
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found, loading environment variables directly")
	}

	// Inisialisasi JWT Secret
	config.InitJwtSecret()

	// Validasi variabel lingkungan
	validateEnv()
}

func validateEnv() {
	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}
}

func main() {
	// Inisialisasi database
	config.InitDB()
	defer config.CloseDB()

	//Setup rute dan jalankan server
	e := routes.SetupRoutes(config.DB)
	e.Logger.Fatal(e.Start(":8080"))
}
