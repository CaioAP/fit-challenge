package usecase

import (
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type AddPersonToChallenge struct {
	PersonRepository    repository.Person
	ChallengeRepository repository.Challenge
}

func (u *AddPersonToChallenge) Execute(personId int, challengeId int) (model.Challenge, error) {
	person, err := u.PersonRepository.FindById(personId)
	if err != nil {
		return model.Challenge{}, err
	}
	challenge, err := u.ChallengeRepository.FindById(challengeId)
	if err != nil {
		return model.Challenge{}, err
	}
	challenge.AddPerson(person)
	err = u.ChallengeRepository.AddPerson(personId, challengeId)
	if err != nil {
		return challenge, err
	}
	return challenge, nil
}
