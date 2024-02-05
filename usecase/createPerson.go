package usecase

import (
	"errors"
	"strings"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type CreatePerson struct {
	PersonRepository *repository.Person
}

func (u *CreatePerson) Execute(input dto.PersonCreate) (model.Person, error) {
	password, err := model.HashPassword(input.Password)
	if err != nil {
		return model.Person{}, err
	}
	person := model.Person{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
	id, err := u.PersonRepository.Create(person, password)
	if err != nil {
		if strings.Contains(err.Error(), "person_email_key") {
			return model.Person{}, errors.New("email already exists")
		}
		if strings.Contains(err.Error(), "person_phone_key") {
			return model.Person{}, errors.New("phone already exists")
		}
		return model.Person{}, err
	}
	person.ID = id
	return person, nil
}
