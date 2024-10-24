package repository

import (
	"database/sql"
	"errors"
	"github.com/valyamoro/internal/domain"
)

type Words struct {
	db *sql.DB
}

func NewWordsRepository(db *sql.DB) *Words {
	return &Words{db}
}

func (w *Words) Create(word domain.Word) (domain.Word, error) {
	_, err := w.db.Exec(
		`INSERT INTO words (
        word, 
        translation_word, 
        dictionary_id, 
        word_language_code, 
        translation_word_language_code
    	) VALUES ($1, $2, $3, $4, $5)`,
		word.Word,
		word.TranslationWord,
		word.DictionaryID,
		word.WordLanguageCode,
		word.TranslationWordLanguageCode,
	)
	if err != nil {
		return domain.Word{}, err
	}

	return word, nil
}

func (w *Words) GetByID(id int64) (domain.Word, error) {
	var word domain.Word
	err := w.db.QueryRow(`SELECT 
    	id, 
    	word,
    	translation_word,
    	dictionary_id,
    	word_language_code,
    	translation_word_language_code
    	FROM words WHERE id=$1`, id).
		Scan(
			&word.ID,
			&word.Word,
			&word.TranslationWord,
			&word.DictionaryID,
			&word.WordLanguageCode,
			&word.TranslationWordLanguageCode,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return word, err
	}

	return word, nil
}

func (w *Words) GetAll() ([]domain.Word, error) {
	rows, err := w.db.Query(`SELECT 
    	id, 
    	word,
    	translation_word,
    	dictionary_id,
    	word_language_code,
    	translation_word_language_code
    	FROM words`)
	if err != nil {
		return nil, err
	}

	words := make([]domain.Word, 0)
	for rows.Next() {
		var word domain.Word
		if err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.TranslationWord,
			&word.DictionaryID,
			&word.WordLanguageCode,
			&word.TranslationWordLanguageCode,
		); err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	return words, rows.Err()
}

func (w *Words) Delete(id int64) (domain.Word, error) {
	word, err := w.GetByID(id)
	if err != nil {
		return domain.Word{}, err
	}

	_, err = w.db.Exec("DELETE FROM words WHERE id=$1", id)

	if err != nil {
		return domain.Word{}, err
	}

	return word, err
}

func (w *Words) Update(id int64, word domain.Word) (domain.Word, error) {
	_, err := w.db.Exec(
		`UPDATE words SET 
                 word=$1,
                 translation_word=$2
                 WHERE id=$3`,
		word.Word,
		word.TranslationWord,
		id,
	)

	if err != nil {
		return domain.Word{}, err
	}

	return w.GetByID(id)
}
