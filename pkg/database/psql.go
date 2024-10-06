package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectionParams struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(params ConnectionParams) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		params.Host,
		params.Port,
		params.Username,
		params.DBName,
		params.SSLMode,
		params.Password,
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
