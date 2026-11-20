package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all runtime configuration loaded from environment variables.
//
// There's no JWT_SECRET / ADMIN_USERNAME / ADMIN_PASSWORD here yet — those
// get added when you build the admin API. See JWT_ADMIN_API_GUIDE.md.
type Config struct {
	AppEnv      string
	Port        string
	DatabaseURL string

	SMTPHost       string
	SMTPPort       string
	SMTPUser       string
	SMTPPassword   string
	ContactToEmail string

	// Medium sync is disabled when MediumUsername is empty.
	MediumUsername     string
	MediumSyncInterval time.Duration
}

// Load reads a .env file (if present) and returns a populated Config.
// Missing optional values fall back to sane defaults so the app can still
// boot in environments where only the essentials are set (e.g. containers
// that inject env vars directly instead of shipping a .env file).
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file found, relying on process environment")
	}

	cfg := &Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "3000"),
		DatabaseURL: getEnv("DATABASE_URL", ""),

		SMTPHost:       getEnv("SMTP_HOST", ""),
		SMTPPort:       getEnv("SMTP_PORT", "587"),
		SMTPUser:       getEnv("SMTP_USER", ""),
		SMTPPassword:   getEnv("SMTP_PASSWORD", ""),
		ContactToEmail: getEnv("CONTACT_TO_EMAIL", ""),

		MediumUsername: getEnv("MEDIUM_USERNAME", ""),
	}

	interval, err := time.ParseDuration(getEnv("MEDIUM_SYNC_INTERVAL", "6h"))
	if err != nil || interval < time.Minute {
		log.Printf("config: invalid MEDIUM_SYNC_INTERVAL, using 6h")
		interval = 6 * time.Hour
	}
	cfg.MediumSyncInterval = interval

	if cfg.DatabaseURL == "" {
		log.Fatal("config: DATABASE_URL is required")
	}

	return cfg
}

func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
