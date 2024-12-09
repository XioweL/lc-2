package main

import (
	"live-code-2-XioweL/config"
)

//var jwtSecret []byte

//func init() {
//	// Muat .env
//	//if err := godotenv.Load(); err != nil {
//	cwd, _ := os.Getwd()
//	envPath := filepath.Join(cwd, "../.env")
//	if err := godotenv.Load(envPath); err != nil {
//		log.Println("No .env file found, loading environment variables directly")
//	}
//
//	// Inisialisasi JWT Secret
//	config.InitJwtSecret()
//
//	// Validasi variabel lingkungan
//	validateEnv()
//}

//func validateEnv() {
//	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
//	for _, v := range requiredVars {
//		if os.Getenv(v) == "" {
//			log.Fatalf("Environment variable %s is not set", v)
//		}
//	}
//}

func main() {
	// Inisialisasi database
	config.InitDB()
	defer config.CloseDB()

	// Setup rute dan jalankan server
	//e := routes.SetupRoutes()
	//e.Logger.Fatal(e.Start(":8080"))
}
