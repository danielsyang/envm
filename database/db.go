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
			name TEXT NOT NULL, 
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	)
	if err != nil {
		log.Fatal(err)
	}

}

func ListEnvironments() []string {
	return []string{"local", "dev", "stage", "prod"}
}
