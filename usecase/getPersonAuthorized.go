package usecase

import (
	"context"

	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/go-chi/jwtauth"
)

type GetPersonAuthorized struct {
	TokenAuth        *jwtauth.JWTAuth
	PersonRepository *repository.Person
}

func (u *GetPersonAuthorized) Execute(token string, ctx context.Context) (model.Person, error) {
	tokenAuth, err := model.DecodeToken(u.TokenAuth, ctx)
	if err != nil {
		return model.Person{}, err
	}
	person, err := u.PersonRepository.FindById(tokenAuth.ID)
	if err != nil {
		return model.Person{}, err
	}
	return person, nil
}
