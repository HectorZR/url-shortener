package migrations

import (
	"github.com/HectorZR/url-shortener/app"
	"gorm.io/gorm"
)

type UrlMigration struct{}

func (um UrlMigration) up(migrator gorm.Migrator) {
	migrator.CreateTable(&app.ShortenedURL{})
}

func (um UrlMigration) down(migrator gorm.Migrator) {
	migrator.DropTable(&app.ShortenedURL{})
}
