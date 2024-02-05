package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/caioap/desafio_bonde/model"
)

type Challenge struct {
	DB *sql.DB
}

func (r *Challenge) Create(challenge model.Challenge) (int, error) {
	var id int
	query := `
		INSERT INTO challenge (name, description, goal, max_per_day, start_date, finish_date, owner_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id
	`
	err := r.DB.QueryRow(
		query,
		challenge.Name,
		challenge.Description,
		challenge.Goal,
		challenge.MaxPerDay,
		challenge.StartDate,
		challenge.FinishDate,
		challenge.Owner.ID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	query = "INSERT INTO person_challenge (person_id, challenge_id) VALUES ($1, $2)"
	_, err = r.DB.Exec(query, challenge.Owner.ID, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Challenge) FindById(id int) (model.Challenge, error) {
	challenge := model.Challenge{}
	query := `
		SELECT 
			c.id,
			c.name,
			c.description,
			c.goal,
			c.max_per_day,
			c.start_date,
			c.finish_date,
			c.owner_id,
			p.name,
			p.email,
			p.phone
		FROM challenge c 
		INNER JOIN person p ON p.id = c.owner_id
		WHERE c.id = $1
	`
	err := r.DB.QueryRow(query, id).Scan(
		&challenge.ID,
		&challenge.Name,
		&challenge.Description,
		&challenge.Goal,
		&challenge.MaxPerDay,
		&challenge.StartDate,
		&challenge.FinishDate,
		&challenge.Owner.ID,
		&challenge.Owner.Name,
		&challenge.Owner.Email,
		&challenge.Owner.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Challenge{}, errors.New("challenge not found")
		}
		return model.Challenge{}, err
	}
	finishDate, _ := time.Parse(time.RFC3339, "2024-04-30T23:59:59-03:00")
	challenge.FinishDate = finishDate
	return challenge, nil
}

func (r *Challenge) FindPeople(id int) ([]model.Person, error) {
	people := []model.Person{}
	query := `
		SELECT p.id, p.name, p.email, p.phone
		FROM person p
		INNER JOIN challenge_person cp ON cp.person_id = p.id
		WHERE cp.challenge_id = $1
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return people, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		person := model.Person{}
		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Email,
			&person.Phone,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

func (r *Challenge) AddPerson(challengeId int, personId int) error {
	query := "INSERT INTO challenge_person (challenge_id, person_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, challengeId, personId)
	if err != nil {
		return err
	}
	return nil
}
