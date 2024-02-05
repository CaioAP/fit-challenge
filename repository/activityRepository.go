package repository

import (
	"database/sql"
	"time"

	"github.com/caioap/desafio_bonde/model"
)

type Activity struct {
	DB *sql.DB
}

func (r *Activity) Create(activity model.Activity) (int, error) {
	var id int
	query := "INSERT INTO activity (type, date, person_id, challenge_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.DB.QueryRow(query, activity.Type, activity.Date, activity.Person.ID, activity.Challenge.ID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Activity) FindByPerson(personId int) ([]model.Activity, error) {
	var activities []model.Activity
	query := "SELECT * FROM activity WHERE person_id = $1"
	rows, err := r.DB.Query(query, personId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ID, &activity.Type, &activity.Date, &activity.Person.ID, &activity.Challenge.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return activities, nil
			}
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func (r *Activity) CountByPersonAndChallengeToday(personId int, challengeId int) (int, error) {
	initialDate := time.Now().Format("2006-01-02") + " 00:00:00"
	finalDate := time.Now().Format("2006-01-02") + " 23:59:59"
	query := `
		SELECT count(*) 
		FROM activity 
		WHERE person_id = $1 
			AND challenge_id = $2 
			AND date >= $3 
			AND date <= $4
	`
	var count int
	err := r.DB.QueryRow(query, personId, challengeId, initialDate, finalDate).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
