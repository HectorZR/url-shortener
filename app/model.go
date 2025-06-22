package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var builder *gorm.DB

func ShortenURL(url string) string {
	// Implement URL shortening logic here
	return "shortened url"
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
