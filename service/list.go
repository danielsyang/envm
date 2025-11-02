package service

import "envm/database"

const SELECT_ALL_ENVIRONMENTS_QUERY = `SELECT name FROM environments;`

func ListEnvironments() ([]string, error) {
	db := database.GetConnection()

	rows, err := db.Query(SELECT_ALL_ENVIRONMENTS_QUERY)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var envs []string

	for rows.Next() {
		var env string
		rows.Scan(&env)

		envs = append(envs, env)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return envs, nil
}
