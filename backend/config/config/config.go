package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}

	Database struct {
		Host string `yaml:"db_host"`
		Port string `yaml:"db_port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}

	JWT struct {
		Secret   string `yaml:"secret"`
		Lifespan string `yaml:"token_lifaspan"`
	}

	Email struct {
		SmtpHOST string `yaml:"smtp_host"`
		SmtpPort string `yaml:"smtp_port"`
		SmtpUser string `yaml:"smtp_user"`
		SmtpPass string `yaml:"smtp_pass"`
	}

	AWS struct {
		Bucket string `yaml:"bucket"`
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	var config Config

	config.Server.Port = os.Getenv("PORT")
	config.Server.Host = os.Getenv("SERVER_HOST")

	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Pass = os.Getenv("DB_PASS")
	config.Database.Name = os.Getenv("DB_NAME")

	config.JWT.Secret = os.Getenv("JWT_SECRET")
	config.JWT.Lifespan = os.Getenv("JWT_TOKEN_LIFESPAN")

	config.Email.SmtpHOST = os.Getenv("SMTP_HOST")
	config.Email.SmtpPort = os.Getenv("SMTP_PORT")
	config.Email.SmtpUser = os.Getenv("SMTP_USER")
	config.Email.SmtpPass = os.Getenv("SMTP_PASS")

	config.AWS.Bucket = os.Getenv("BUCKET")

	return &config, nil
}
