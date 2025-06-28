package app

import (
	"fmt"
	"os"
)

func GetEnvVars() map[string]string {
	envVars := make(map[string]string)
	envVars["DB_HOST"] = os.Getenv("DB_HOST")
	envVars["DB_PORT"] = os.Getenv("DB_PORT")
	envVars["DB_USER"] = os.Getenv("DB_USER")
	envVars["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVars["DB_NAME"] = os.Getenv("DB_NAME")
	return envVars
}

func GetPostgresDSN() string {
	envVars := GetEnvVars()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York", envVars["DB_HOST"], envVars["DB_USER"], envVars["DB_PASSWORD"], envVars["DB_NAME"], envVars["DB_PORT"])
}
