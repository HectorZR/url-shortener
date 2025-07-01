package shared

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const base62chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
 * GetEnvVars returns a map of environment variables.
 */
func GetEnvVars() map[string]string {
	envVars := make(map[string]string)
	envVars["DB_HOST"] = os.Getenv("DB_HOST")
	envVars["DB_PORT"] = os.Getenv("DB_PORT")
	envVars["DB_USER"] = os.Getenv("DB_USER")
	envVars["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVars["DB_NAME"] = os.Getenv("DB_NAME")
	return envVars
}

/*
 * GetPostgresDSN returns a PostgreSQL data source name.
 */
func GetPostgresDSN() string {
	envVars := GetEnvVars()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", envVars["DB_HOST"], envVars["DB_USER"], envVars["DB_PASSWORD"], envVars["DB_NAME"], envVars["DB_PORT"])
}

/*
 * GenerateCodeFromHash returns a hash code from a URL.
 */
func GenerateCodeFromHash(url string) string {
	hash := sha256.Sum256([]byte(url))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

/*
 * InitDB initializes the database.
 */
func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(GetPostgresDSN()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

/*
 * EncodeBase62 encodes a number to a base62 string.
 */
func EncodeBase62(num uint) string {
	if num == 0 {
		return string(base62chars[0])
	}

	result := ""

	for num > 0 {
		result = string(base62chars[num%62]) + result
		num /= 62
	}

	return result
}

/*
 * DecodeBase62 decodes a base62 string to a number.
 */
func DecodeBase62(str string) uint {
	num := uint(0)

	for _, char := range str {
		num *= 62
		for i, b := range base62chars {
			if b == char {
				num += uint(i)
				break
			}
		}
	}

	return num
}
