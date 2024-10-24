package service

import "github.com/valyamoro/internal/domain"

type DictionariesRepository interface {
	Create(dictionary domain.Dictionary) (domain.Dictionary, error)
	GetByID(id int64) (domain.Dictionary, error)
	GetAll() ([]domain.Dictionary, error)
	Delete(id int64) (domain.Dictionary, error)
	Update(id int64, inp domain.Dictionary) (domain.Dictionary, error)
}

type Dictionaries struct {
	repo DictionariesRepository
}

func NewDictionariesService(repo DictionariesRepository) *Dictionaries {
	return &Dictionaries{
		repo: repo,
	}
}

func (d *Dictionaries) Create(dictionary domain.Dictionary) (domain.Dictionary, error) {
	return d.repo.Create(dictionary)
}

func (d *Dictionaries) GetByID(id int64) (domain.Dictionary, error) {
	return d.repo.GetByID(id)
}

func (d *Dictionaries) GetAll() ([]domain.Dictionary, error) {
	return d.repo.GetAll()
}

func (d *Dictionaries) Delete(id int64) (domain.Dictionary, error) {
	return d.repo.Delete(id)
}

func (d *Dictionaries) Update(id int64, dictionary domain.Dictionary) (domain.Dictionary, error) {
	return d.repo.Update(id, dictionary)
}
