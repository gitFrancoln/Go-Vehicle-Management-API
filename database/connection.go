package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {

		dir, derr := os.Getwd()
		if derr != nil {
			return nil, fmt.Errorf("error obteniendo working dir: %w", derr)
		}
		for {
			envPath := filepath.Join(dir, ".env")
			if _, statErr := os.Stat(envPath); statErr == nil {
				if loadErr := godotenv.Load(envPath); loadErr != nil {
					return nil, fmt.Errorf("error loading .env from %s: %w", envPath, loadErr)
				}
				break
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				return nil, fmt.Errorf(".env not found in working dir or parent directories: %w", err)
			}
			dir = parent
		}
	}

	// Construir DSN desde variables de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying SQL DB: %w", err)
	}
	_ = sqlDB
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying SQL DB for close: %w", err)
	}
	return sqlDB.Close()
}
