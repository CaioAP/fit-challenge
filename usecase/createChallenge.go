package usecase

import (
	"errors"
	"time"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type CreateChallenge struct {
	PersonRepository    *repository.Person
	ChallengeRepository *repository.Challenge
	CreateRanking       *CreateRanking
}

func (u *CreateChallenge) Execute(input dto.ChallengeCreate) (model.Challenge, error) {
	person, err := u.PersonRepository.FindById(input.PersonId)
	if err != nil {
		return model.Challenge{}, err
	}
	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return model.Challenge{}, errors.New("invalid start date")
	}
	finishDate, err := time.Parse(time.RFC3339, input.FinishDate)
	if err != nil {
		return model.Challenge{}, errors.New("invalid finish date")
	}
	challenge := model.Challenge{
		Name:        input.Name,
		Description: input.Description,
		Goal:        input.Goal,
		MaxPerDay:   input.MaxPerDay,
		StartDate:   startDate,
		FinishDate:  finishDate,
		CreatedAt:   time.Now(),
		Owner:       person,
		Person:      []model.Person{person},
	}
	id, err := u.ChallengeRepository.Create(challenge)
	if err != nil {
		return model.Challenge{}, err
	}
	challenge.ID = id
	go u.CreateRanking.Execute(person, challenge)
	return challenge, nil
}
