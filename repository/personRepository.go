package repository

import (
	"database/sql"
	"errors"

	"github.com/caioap/desafio_bonde/model"
)

type Person struct {
	DB *sql.DB
}

func (r *Person) Create(person model.Person, password string) (int, error) {
	var id int
	query := "INSERT INTO person (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.DB.QueryRow(query, person.Name, person.Email, person.Phone, password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Person) FindById(id int) (model.Person, error) {
	person := model.Person{}
	query := "SELECT id, name, phone, email, created_at FROM person WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(
		&person.ID,
		&person.Name,
		&person.Phone,
		&person.Email,
		&person.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Person{}, errors.New("person not found")
		}
		return model.Person{}, err
	}
	return person, nil
}
