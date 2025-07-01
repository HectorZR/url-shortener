package shortener

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/HectorZR/url-shortener/shared"
	"gorm.io/gorm"
)

/*
 * ShortenedURL model
 */
type ShortenedURL struct {
	gorm.Model
	OriginalURL string `gorm:"not null;unique"`
}

func ShortenURL(url string, db *gorm.DB) ShortenedURL {
	var existingUrl ShortenedURL
	db.First(&existingUrl, "original_url = ?", url)

	if existingUrl.ID != 0 {
		return existingUrl
	}
	newUrl := ShortenedURL{OriginalURL: url}
	db.Create(&newUrl)
	return newUrl
}

func GetOriginalURL(shortCode string, db *gorm.DB) (ShortenedURL, error) {
	id := shared.DecodeBase62(shortCode)
	var shortened ShortenedURL
	db.First(&shortened, id)

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
