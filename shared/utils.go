package shared

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	BASE_62_CHARS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DEV           = "dev"
	PROD          = "prod"
)

// Private structure for configuration
type config struct {
	DBHost                string
	DBPort                string
	DBUser                string
	DBPassword            string
	DBName                string
	Port                  string
	PathPrefix            string
	Env                   string
	SiteKey               string
	SecretKey             string
	ProjectID             string
	CaptchaAction         string
	GoogleCredentialsJson string
}

/*
 * GetEnvVars returns a struct of environment variables.
 */
func GetEnvVars() config {
	c := config{}
	c.Port = os.Getenv("PORT")
	c.Env = os.Getenv("ENV")
	c.PathPrefix = os.Getenv("PATH_PREFIX")
	c.DBHost = os.Getenv("DB_HOST")
	c.DBPort = os.Getenv("DB_PORT")
	c.DBUser = os.Getenv("DB_USER")
	c.DBName = os.Getenv("DB_NAME")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.SiteKey = os.Getenv("SITE_KEY")
	c.SecretKey = os.Getenv("SECRET_KEY")
	c.ProjectID = os.Getenv("PROJECT_ID")
	c.CaptchaAction = os.Getenv("CAPTCHA_ACTION")
	c.GoogleCredentialsJson = os.Getenv("GOOGLE_CREDENTIALS_JSON")
	return c
}

/*
 * GetPostgresDSN returns a PostgreSQL data source name.
 */
func GetPostgresDSN() string {
	configs := GetEnvVars()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", configs.DBHost, configs.DBUser, configs.DBPassword, configs.DBName, configs.DBPort)
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
		return string(BASE_62_CHARS[0])
	}

	result := ""

	for num > 0 {
		result = string(BASE_62_CHARS[num%62]) + result
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
		for i, b := range BASE_62_CHARS {
			if b == char {
				num += uint(i)
				break
			}
		}
	}

	return num
}
