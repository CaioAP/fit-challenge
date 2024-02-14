package usecase

import (
	"errors"
	"strings"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/go-chi/jwtauth"
)

type Register struct {
	TokenAuth        *jwtauth.JWTAuth
	PersonRepository *repository.Person
}

func (u *Register) Execute(input dto.Register) (model.AuthOutput, error) {
	password, err := model.HashPassword(input.Password)
	if err != nil {
		return model.AuthOutput{}, err
	}
	person := model.Person{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
	id, err := u.PersonRepository.Create(person, password)
	if err != nil {
		if strings.Contains(err.Error(), "person_email_key") {
			return model.AuthOutput{}, errors.New("email already exists")
		}
		if strings.Contains(err.Error(), "person_phone_key") {
			return model.AuthOutput{}, errors.New("phone already exists")
		}
		return model.AuthOutput{}, err
	}
	person.ID = id
	token := model.CreateToken(u.TokenAuth, person.ID, person.Name)
	return model.AuthOutput{
		Token: token,
		ID:    person.ID,
		Name:  person.Name,
	}, nil
}
