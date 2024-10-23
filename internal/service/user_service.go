package service

import "github.com/valyamoro/internal/domain"

type UsersRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	GetByID(id int) (domain.User, error)
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

func (u *Users) GetByUsername(username string) (domain.User, error) {
	return u.repo.GetByUsername(username)
}

func (u *Users) GetByID(id int) (domain.User, error) {
	return u.repo.GetByID(id)
}
