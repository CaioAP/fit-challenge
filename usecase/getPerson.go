package usecase

import (
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type GetPerson struct {
	PersonRepository *repository.Person
}

func (u *GetPerson) Execute(id int) (model.Person, error) {
	person, err := u.PersonRepository.FindById(id)
	if err != nil {
		return model.Person{}, err
	}
	return person, nil
}
