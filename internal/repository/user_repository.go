package repository

import (
	"database/sql"
	"github.com/valyamoro/internal/domain"
	"github.com/valyamoro/internal/auth"
	"errors"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u *Users) Create(user domain.User) (domain.User, error) {
	_, err := u.db.Exec(
		"INSERT INTO users (username, password) values ($1, $2)",
		user.Username,
		auth.HashPassword(user.Password),
	)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *Users) GetByUsername(username string) (domain.User, error) {
	var user domain.User 
	err := u.db.QueryRow(`SELECT id, username, password FROM users WHERE username=$1`, username).
		Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return user, err 
	}

	return user, nil
}
