package database

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("sqlite", "./envm.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS environments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	)
	if err != nil {
		log.Fatal(err)
	}

}

func GetConnection() *sql.DB {
	if db == nil {
		log.Fatal("database not initialized")
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal("database connection lost: %w", err)
	}

	return db
}
