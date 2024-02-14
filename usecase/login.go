package usecase

import (
	"errors"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/go-chi/jwtauth"
)

type Login struct {
	TokenAuth        *jwtauth.JWTAuth
	PersonRepository *repository.Person
}

func (u *Login) Execute(dto dto.Login) (model.AuthOutput, error) {
	person, err := u.PersonRepository.FindByEmail(dto.Email)
	if err != nil {
		return model.AuthOutput{}, err
	}
	if !model.CheckPasswordHash(dto.Password, person.Password) {
		return model.AuthOutput{}, errors.New("password does not match")
	}
	token := model.CreateToken(u.TokenAuth, person.ID, person.Name)
	return model.AuthOutput{
		Token: token,
		ID:    person.ID,
		Name:  person.Name,
	}, nil
}
