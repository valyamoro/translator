package repository

import (
	"database/sql"
	"errors"
	"github.com/valyamoro/internal/domain"
)

type Dictionaries struct {
	db *sql.DB
}

func NewDictionaryRepository(db *sql.DB) *Dictionaries {
	return &Dictionaries{db}
}

func (d *Dictionaries) Create(dictionary domain.Dictionary) (domain.Dictionary, error) {
	_, err := d.db.Exec(
		"INSERT INTO dictionaries (name, description, user_id) values ($1, $2, $3)",
		dictionary.Name,
		dictionary.Description,
		dictionary.UserID,
	)
	if err != nil {
		return domain.Dictionary{}, err
	}

	return dictionary, nil
}

func (d *Dictionaries) GetByID(id int64) (domain.Dictionary, error) {
	var dictionary domain.Dictionary
	err := d.db.QueryRow("SELECT id, name, user_id FROM dictionaries WHERE id=$1", id).
		Scan(
			&dictionary.ID,
			&dictionary.Name,
			&dictionary.Description,
			&dictionary.UserID,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return dictionary, err
	}

	return dictionary, nil
}

func (d *Dictionaries) GetAll() ([]domain.Dictionary, error) {
	rows, err := d.db.Query("SELECT id, name, description, user_id FROM dictionaries")
	if err != nil {
		return nil, err
	}

	dictionaries := make([]domain.Dictionary, 0)
	for rows.Next() {
		var dictionary domain.Dictionary
		if err := rows.Scan(
			&dictionary.ID,
			&dictionary.Name,
			&dictionary.Description,
			&dictionary.UserID,
		); err != nil {
			return nil, err
		}

		dictionaries = append(dictionaries, dictionary)
	}

	return dictionaries, rows.Err()
}

func (d *Dictionaries) Delete(id int64) (domain.Dictionary, error) {
	dictionary, err := d.GetByID(id)
	if err != nil {
		return domain.Dictionary{}, err
	}

	_, err = d.db.Exec("DELETE FROM dictionaries WHERE id=$1", id)

	if err != nil {
		return domain.Dictionary{}, err
	}

	return dictionary, err
}

func (d *Dictionaries) Update(id int64, inp domain.UpdateDictionaryInput) (domain.Dictionary, error) {
	_, err := d.db.Exec(
		"UPDATE dictionaries SET name=$1, description=$2, user_id=$3 WHERE id=$4",
		inp.Name,
		inp.Description,
		inp.UserID,
		id,
	)
	if err != nil {
		return domain.Dictionary{}, err
	}

	return d.GetByID(id)
}
