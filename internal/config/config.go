package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabaseUser string
	DatabasePass string
	DatabaseHost string
	DatabasePort string
	DatabaseName string
}

var Env *Config

func LoadConfig() {

	rootDir := findProjectRoot()
	envPath := filepath.Join(rootDir, ".env")

	// Cargar el archivo .env
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Env = &Config{
		Port:         getEnvOrDefault("PORT", "8080"),
		DatabaseUser: getEnvOrDefault("DB_USER", "root"),
		DatabasePass: getEnvOrDefault("DB_PASSWORD", "7777"),
		DatabaseHost: getEnvOrDefault("DB_HOST", "127.0.0.1"),
		DatabasePort: getEnvOrDefault("DB_PORT", "3306"),
		DatabaseName: getEnvOrDefault("DB_NAME", "tpirso"),
	}

	log.Printf("Config loaded successfully: Port=%s, DB=%s:%s",
		Env.Port, Env.DatabaseHost, Env.DatabasePort)
}

func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "."
		}
		dir = parent
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
