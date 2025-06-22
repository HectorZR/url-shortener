package app

import (
	"database/sql"
	"fmt"
)

type AppModel struct {
	db *sql.DB
}

func (am *AppModel) InitDB() {
	db, err := sql.Open("sqlite3", "./url_shortener.db")
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}

	am.db = db
}

func ShortenURL(url string) string {
	// Implement URL shortening logic here
	return "shortened url"
}
