package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *Challenge) FindByPerson(personId int) ([]model.Challenge, error) {
	challenges := []model.Challenge{}
	queryParticipants := "select count(*) from person_challenge pc where challenge_id = c.id"
	queryRanking := `
		select row_number() over() as "ranking" 
		from (
			select r.person_id 
			from ranking r
			where r.challenge_id = c.id
			order by r.completed desc, r.updated_at asc
		) t
		where t.person_id = cp.person_id
	`
	queryActivities := `
		select count(*) 
		from activity a
		where a.challenge_id = c.id and a.person_id = cp.person_id 
	`
	query := fmt.Sprintf(`
		SELECT c.*, (%s) as participants, (%s) as ranking, (%s) as activites
		FROM challenge c
		INNER JOIN person_challenge cp ON cp.challenge_id = c.id
		WHERE cp.person_id = $1;
	`, queryParticipants, queryRanking, queryActivities)
	rows, err := r.DB.Query(query, personId)
	if err != nil {
		if err == sql.ErrNoRows {
			return challenges, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		challenge := model.Challenge{}
		err := rows.Scan(
			&challenge.ID,
			&challenge.Name,
			&challenge.Description,
			&challenge.Goal,
			&challenge.MaxPerDay,
			&challenge.StartDate,
			&challenge.FinishDate,
			&challenge.Owner.ID,
			&challenge.CreatedAt,
			&challenge.Participants,
			&challenge.Ranking,
			&challenge.Activities,
		)
		if err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	return challenges, nil
}

func (r *Challenge) AddPerson(challengeId int, personId int) error {
	query := "INSERT INTO challenge_person (challenge_id, person_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, challengeId, personId)
	if err != nil {
		return err
	}
	return nil
}
