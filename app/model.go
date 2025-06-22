package app

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
	"regexp"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortenedURL struct {
	gorm.Model
	OriginalURL string `gorm:"not null;unique"`
	ShortURL    string `gorm:"not null"`
}

func ShortenURL(url string) string {
	db := initDB()

	shortUrl := generateCodeFromHash(url)

	shortened := ShortenedURL{
		OriginalURL: url,
		ShortURL:    shortUrl,
	}

	db.FirstOrCreate(&shortened)

	return shortened.ShortURL
}

func GetOriginalURL(shortURL string) (string, error) {
	db := initDB()

	shortened := ShortenedURL{}
	db.First(&shortened, "short_url = ?", shortURL)

	if shortened.ID == 0 {
		return "", errors.New("Short URL not found")
	}

	return shortened.OriginalURL, nil
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

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("url_shortener.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&ShortenedURL{})

	return db
}

func generateCodeFromHash(url string) string {
	hash := sha256.Sum256([]byte(url))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
