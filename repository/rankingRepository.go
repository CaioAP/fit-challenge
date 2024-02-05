package repository

import (
	"database/sql"
	"errors"

	"github.com/caioap/desafio_bonde/model"
)

type Ranking struct {
	DB *sql.DB
}

func (r *Ranking) Create(ranking model.Ranking) (int, error) {
	query := `INSERT INTO ranking (completed, challenge_id, person_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.DB.QueryRow(query, ranking.Completed, ranking.Challenge.ID, ranking.Person.ID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Ranking) FindOne(challengeId int, personId int) (model.Ranking, error) {
	query := `
		SELECT id, completed, updated_at, challenge_id, person_id
		FROM ranking 
		WHERE challenge_id = $1
			AND person_id = $2
	`
	ranking := model.Ranking{}
	err := r.DB.QueryRow(query, challengeId, personId).Scan(
		&ranking.ID,
		&ranking.Completed,
		&ranking.UpdatedAt,
		&ranking.Challenge.ID,
		&ranking.Person.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Ranking{}, errors.New("ranking not found")
		}
		return model.Ranking{}, err
	}
	return ranking, nil
}

func (r *Ranking) FindByChallenge(challengeId int) ([]model.Ranking, error) {
	query := `
		SELECT id, completed, updated_at, challenge_id, person_id 
		FROM ranking 
		WHERE challenge_id = $1
		ORDER BY completes DESC, updated_at ASC
	`
	rows, err := r.DB.Query(query, challengeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rankings := []model.Ranking{}
	for rows.Next() {
		var ranking model.Ranking
		err := rows.Scan(
			&ranking.ID,
			&ranking.Completed,
			&ranking.UpdatedAt,
			&ranking.Challenge.ID,
			&ranking.Person.ID,
		)
		if err != nil {
			return rankings, err
		}
		rankings = append(rankings, ranking)
	}
	return rankings, nil
}

func (r *Ranking) Update(ranking model.Ranking) (model.Ranking, error) {
	query := `
		UPDATE ranking SET completed = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.DB.Exec(query, ranking.Completed, ranking.UpdatedAt, ranking.ID)
	if err != nil {
		return model.Ranking{}, err
	}
	return ranking, nil
}
