package service

import (
	"database/sql"
	"envm/database"
	"fmt"
)

const SELECT_ENVIRONMENT_CONTENT_QUERY = `SELECT content FROM environments WHERE name = $1;`

func GenerateEnvironment(env string) (string, error) {
	db := database.GetConnection()

	var content string

	err := db.QueryRow(SELECT_ENVIRONMENT_CONTENT_QUERY, env).Scan(&content)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no environment found for %s", env)
	}

	if err != nil {
		return "", err
	}

	return content, nil
}
