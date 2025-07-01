package main

import (
	"fmt"
	"os"

	"github.com/HectorZR/url-shortener/migrations"
	"github.com/HectorZR/url-shortener/shared"
)

func main() {
	if len(os.Args) <= 2 {
		invalidCommand()
		help()
		return
	}

	order := os.Args[1]
	action := migrations.AllowedDirection(os.Args[2])

	switch {
	case order == migrations.MIGRATE && (action == migrations.UP || action == migrations.DOWN):
		db := shared.InitDB()
		migrations.Handler(action, db.Migrator())

	case order == migrations.MIGRATE && action == migrations.RESET:
		db := shared.InitDB()
		migrations.Handler(migrations.DOWN, db.Migrator())
		migrations.Handler(migrations.UP, db.Migrator())

	default:
		invalidCommand()
		help()
	}
}

func help() {
	fmt.Println(`Available commands:
	migrate
		up    - Run all migrations up
		down  - Run all migrations down
		reset - Reset all migrations

Tip: if you are using Docker, don't forget to run the migrations inside the container.
`)
}

func invalidCommand() {
	fmt.Println("\033[31mInvalid command\033[0m")
}
