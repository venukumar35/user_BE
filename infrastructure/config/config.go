package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Name            string
	Port            string
	DbDsn           string
	ProdutDb        string
	TcpPort         string
	JwtSecretKey    string
	TokenDuration   time.Duration
	SmtpHost        string
	SmtpPort        string
	SmtpUserName    string
	SmtpPassword    string
	SmtpDisplayName string
	SyncData        bool
	SyncTime        int16
}

var Config *Configuration

func Load() error {

	_ = godotenv.Load()

	Config = &Configuration{
		Name:            getEnvOrError("PROJECT_NAME"),
		Port:            getEnvOrError("PORT"),
		DbDsn:           getEnvOrError("DATABASE_URL"),
		ProdutDb:        getEnvOrError("PRODUCT_DATABASE_URL"),
		JwtSecretKey:    getEnvOrError("SECRET_KEY"),
		TokenDuration:   time.Hour * 24,
		SmtpHost:        getEnvOrError("SMTP_HOST"),
		SmtpPort:        getEnvOrError("SMTP_PORT"),
		SmtpUserName:    getEnvOrError("SMTP_USERNAME"),
		SmtpDisplayName: getEnvOrError("SMTP_DISPLAY_NAME"),
		SmtpPassword:    getEnvOrError("SMTP_PASSWORD"),
		SyncData:        bool(getEnvAsBool("SYNC_DATA")),
		SyncTime:        int16(getEnvAsInt("SYNC_TIME")),
	}

	return nil
}

func getEnvOrError(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s not set", key))
}

func getEnvAsInt(key string) int64 {
	valueStr := getEnvOrError(key)
	var value int64
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		log.Printf("\nError loading %s: %v", key, err)
		panic(err)
	}
	return value
}

func getEnvAsBool(key string) bool {
	valueStr := getEnvOrError(key)
	var value bool
	_, err := fmt.Sscanf(valueStr, "%t", &value)
	if err != nil {
		log.Printf("\nError loading %s: %v", key, err)
		panic(err)
	}
	return value
}
