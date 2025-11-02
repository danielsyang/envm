package service

import (
	"envm/database"
	"strings"
)

const INSERT_ENVIRONMENT_QUERY = `INSERT INTO environments (name, content) VALUES ($1, $2);`

const DELETE_ENVIRONMENT_QUERY = `DELETE FROM environments WHERE name = $1;`

func deleteEnvironment(env string) error {
	db := database.GetConnection()

	_, err := db.Exec(DELETE_ENVIRONMENT_QUERY, env)

	return err
}

func insertEnvironment(env string, file []byte) error {
	db := database.GetConnection()

	_, err := db.Exec(INSERT_ENVIRONMENT_QUERY, env, string(file))

	return err
}

func AddEnvironment(env string, file []byte) error {

	err := insertEnvironment(env, file)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			err = deleteEnvironment(env)

			if err != nil {
				return err
			}

			err = insertEnvironment(env, file)

			if err != nil {
				return err
			}

		}

		return err
	}

	return nil
}
