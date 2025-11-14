package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading .env file...")
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Println("⚠️ Error loading .env:", loadErr)
		}
	} else {
		log.Println(".env not found — skipping (Render/Railway mode)")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Println("⚠️ WARNING: Missing one or more environment variables for database connection")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
	)

	log.Println("Connecting to DB:", host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ Error connecting to DB: %w", err)
	}

	log.Println("✅ Connected to database successfully!")
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting SQL DB for close: %w", err)
	}
	return sqlDB.Close()
}
