package migrations

import (
	"github.com/HectorZR/url-shortener/modules/shortener"
	"gorm.io/gorm"
)

type UrlMigration struct{}

func (um UrlMigration) up(migrator gorm.Migrator) {
	migrator.CreateTable(&shortener.ShortenedURL{})
}

func (um UrlMigration) down(migrator gorm.Migrator) {
	migrator.DropTable(&shortener.ShortenedURL{})
}
