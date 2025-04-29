package initializers

import (
 "fmt"
 "log"
 "os"
"time"
 "gorm.io/driver/postgres"
 "gorm.io/gorm"

 "github.com/roqiaahmed/wikidocify/pkg/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"postgres://%s:%s@db:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
		if err == nil {
			log.Println("Connected to the database.")
			DB.AutoMigrate(&models.Document{}) 
			return
		}
		log.Println("Retrying DB connection...")
		time.Sleep(2 * time.Second)

	log.Fatalf("Failed to connect to database: %v", err)
}

