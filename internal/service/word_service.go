package service

import "github.com/valyamoro/internal/domain"

type WordsRepository interface {
	Create(word domain.Word) (domain.Word, error)
	GetByID(id int64) (domain.Word, error)
	GetAll() ([]domain.Word, error)
	Delete(id int64) (domain.Word, error)
	Update(id int64, word domain.Word) (domain.Word, error)
}

type Words struct {
	repo WordsRepository
}

func NewWordsService(repo WordsRepository) *Words {
	return &Words{
		repo: repo,
	}
}

func (w *Words) Create(word domain.Word) (domain.Word, error) {
	return w.repo.Create(word)
}

func (w *Words) GetByID(id int64) (domain.Word, error) {
	return w.repo.GetByID(id)
}

func (w *Words) GetAll() ([]domain.Word, error) {
	return w.repo.GetAll()
}

func (w *Words) Delete(id int64) (domain.Word, error) {
	return w.repo.Delete(id)
}

func (w *Words) Update(id int64, word domain.Word) (domain.Word, error) {
	return w.repo.Update(id, word)
}
