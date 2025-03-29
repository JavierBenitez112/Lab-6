package config

import (
	"fmt"
	"os"
	"time"

	"laliga-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Intento %d/%d de conexión a la base de datos falló: %v\n", i+1, maxRetries, err)
		if i < maxRetries-1 {
			fmt.Printf("Reintentando en %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return fmt.Errorf("error al conectar con la base de datos después de varios intentos: %v", err)
	}

	// Auto-migrar el modelo Match
	if err = DB.AutoMigrate(&models.Match{}); err != nil {
		return fmt.Errorf("error al migrar la base de datos: %v", err)
	}

	return nil
}
