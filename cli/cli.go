package main

import (
	"fmt"
	"os"

	"github.com/HectorZR/url-shortener/app"
	"github.com/HectorZR/url-shortener/migrations"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid command")
		help()
		return
	}

	order := os.Args[1]
	action := os.Args[2]

	switch {
	case order == migrations.MIGRATE && (action == migrations.UP || action == migrations.DOWN):
		db := app.InitDB()
		migrations.Handler(action, db.Migrator())
	default:
		fmt.Println("Invalid command")
		help()
	}
}

func help() {
	fmt.Println(`Available commands:
	migrate
		up    - Run all migrations up
		down  - Run all migrations down

Tip: if you are using Docker, don't forget to run the migrations inside the container.
`)
}
