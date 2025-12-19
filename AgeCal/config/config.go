package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost5432/userdb?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
