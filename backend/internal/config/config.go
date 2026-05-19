package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	AppPort        string
	AppURL         string
	AppPrefix      string
	CorsAllows     string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	JWTSecret      string
	JWTExpiresIn   string
	AdminEmail     string
	AdminPassword  string
	AdminFirstName string
	AdminLastName  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variable")
	}

	config := &Config{
		// ค่าปลอดภัย
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "8080"),
		AppURL:       getEnv("APP_URL", "http://localhost:8080"),
		AppPrefix:    getEnv("APP_PREFIX", "/queue-doc-api"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5431"),
		DBUser:       getEnv("DB_USER", "cmulifelonged"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
		CorsAllows:   getEnv("CORS_ALLOWED_ORIGINS", "https://www.lifelong.cmu.ac.th"),

		// ค่าไม่ปลอดภัย
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		AdminEmail:    getEnv("ADMIN_EMAIL", ""),
		AdminPassword: getEnv("ADMIN_PASSWORD", "1234"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	if config.AppEnv == "production" {
		if config.DBPassword == "" {
			return fmt.Errorf("DB_PASS is required for production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET is required for production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWT_SECRET must be at least 32 charecters long for production")
		}
		if config.DBSSLMode == "disable" {
			log.Println("Warnig: SSL is disabled for database connection in production")
		}
		if config.AdminEmail == "" {
			return fmt.Errorf("ADMIN_EMAIL is required for production environment")
		}
	}

	if config.AdminEmail != "" && !isValidEmail(config.AdminEmail) {
		return errors.New("ADMIN_EMAIL must be a valid email address")
	}

	if config.DBName == "" {
		return fmt.Errorf("DB_Name is required for production environment")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ฟังก์ชันตรวจสอบอีเมลว่าถูกต้องหรือไม่
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// ตรวจสอบพื้นฐาน - ต้องมี @ และ . และไม่เริ่มหรือจบด้วย @
	return len(email) > 0 &&
		len(email) <= 254 &&
		strings.Contains(email, "@") &&
		strings.Contains(email, ".") &&
		!strings.HasPrefix(email, "@") &&
		!strings.HasSuffix(email, "@")
}
