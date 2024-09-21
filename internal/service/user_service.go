package service

import "github.com/valyamoro/internal/domain"

type UsersRepository interface {
	Create(user domain.User) (domain.User, error)
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
