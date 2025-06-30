package app

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

/*
 * ShortenedURL model
 */
type ShortenedURL struct {
	gorm.Model
	OriginalURL string `gorm:"not null;unique"`
	ShortURL    string `gorm:"not null"`
}

func ShortenURL(url string, db *gorm.DB) ShortenedURL {
	shortUrl := GenerateCodeFromHash(url)

	existingUrl, err := GetOriginalURL(shortUrl, db)
	if err != nil {
		existingUrl.OriginalURL = url
		existingUrl.ShortURL = shortUrl
		db.Create(&existingUrl)
	}

	return existingUrl
}

func GetOriginalURL(shortURL string, db *gorm.DB) (ShortenedURL, error) {
	shortened := ShortenedURL{}
	db.First(&shortened, "short_url = ?", shortURL)

	if shortened.ID == 0 {
		return ShortenedURL{}, errors.New("Short URL not found")
	}

	return shortened, nil
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
