package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

type AllowedDirection string

const (
	MIGRATE                  = "migrate"
	RESET                    = "reset"
	UP      AllowedDirection = "up"
	DOWN    AllowedDirection = "down"
)

type IMigration interface {
	up(migrator gorm.Migrator)
	down(migrator gorm.Migrator)
}

func Handler(action AllowedDirection, migrator gorm.Migrator) {
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
