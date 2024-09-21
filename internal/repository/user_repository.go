package repository

import (
	"database/sql"
	"github.com/valyamoro/internal/domain"
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
		user.Password,
	)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
