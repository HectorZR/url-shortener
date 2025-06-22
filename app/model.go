package app

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var builder *gorm.DB

func ShortenURL(url string) string {
	// Implement URL shortening logic here
	return "shortened url"
}

func ValidateURL(u string) error {
	if u == "" {
		return errors.New("URL cannot be empty")
	}

	if strings.Contains(u, " ") {
		return errors.New("URL cannot contain spaces")
	}

	_, err := url.ParseRequestURI(u)
	if err != nil {
		return errors.New("URL is not valid")
	}

	re := regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	if !re.MatchString(u) {
		return errors.New("URL format is invalid")
	}

	return nil
}

type ShortenedURL struct {
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ShortURL    string `gorm:"not null"`
}

func InitDB() {
	db, err := gorm.Open(sqlite.Open("url_shortener.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ShortenedURL{})

	builder = db
}
