package usecase

import (
	"errors"
	"time"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type CreateActivity struct {
	PersonRepository    *repository.Person
	ChallengeRepository *repository.Challenge
	ActivityRepository  *repository.Activity
	UpdateRanking       *UpdateRanking
}

func (u *CreateActivity) Execute(input dto.ActivityCreate) (model.Activity, error) {
	person, err := u.PersonRepository.FindById(input.PersonId)
	if err != nil {
		return model.Activity{}, err
	}
	challenge, err := u.ChallengeRepository.FindById(input.ChallengeId)
	if err != nil {
		return model.Activity{}, err
	}
	activitiesToday, err := u.ActivityRepository.CountByPersonAndChallengeToday(input.PersonId, input.ChallengeId)
	if err != nil {
		return model.Activity{}, err
	}
	if !challenge.ValidateMaxPerDay(activitiesToday) {
		return model.Activity{}, errors.New("max activities per day reached")
	}
	date := time.Now()
	if input.Date != "" {
		date, err = time.Parse(time.RFC3339, input.Date)
		if err != nil {
			return model.Activity{}, errors.New("invalid date")
		}
	}
	if !challenge.ValidateDate(date) {
		return model.Activity{}, errors.New("challenge is not active")
	}
	activity := model.Activity{
		Type:      input.Type,
		Date:      time.Now(),
		Person:    person,
		Challenge: challenge,
	}
	id, err := u.ActivityRepository.Create(activity)
	if err != nil {
		return model.Activity{}, err
	}
	activity.ID = id
	go u.UpdateRanking.Execute(person.ID, challenge.ID)
	return activity, nil
}
