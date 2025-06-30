package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	MIGRATE = "migrate"
	UP      = "up"
	DOWN    = "down"
)

type IMigration interface {
	up(migrator gorm.Migrator)
	down(migrator gorm.Migrator)
}

func Handler(action string, migrator gorm.Migrator) {
	migrations := []IMigration{
		&UrlMigration{},
	}

	for _, migration := range migrations {
		if action == UP {
			migration.up(migrator)
		} else if action == DOWN {
			migration.down(migrator)
		}
	}

	fmt.Println("\033[32mAll migrations ran successfully!\033[0m")
}
