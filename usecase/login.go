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

func (u *Login) Execute(dto dto.Login) (string, error) {
	person, err := u.PersonRepository.FindByEmail(dto.Email)
	if err != nil {
		return "", err
	}
	if !model.CheckPasswordHash(dto.Password, person.Password) {
		return "", errors.New("invalid password")
	}
	token := model.CreateToken(u.TokenAuth, person.ID)
	return token, nil
}
