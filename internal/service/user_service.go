package service

import "github.com/valyamoro/internal/domain"
import "github.com/valyamoro/pkg/errors"
import "github.com/valyamoro/internal/auth"

type UsersRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
}

type Users struct {
	repo UsersRepository
}

func NewUsersService(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) Create(user domain.User) (domain.User, error) {
	return u.repo.Create(user)
}

func (u *Users) Authenticate(username, password string) (bool, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return false, errors.ErrInvalidCredentials
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return false, errors.ErrInvalidCredentials
	}

	return true, nil
}
