package usecase

import (
	"context"

	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/go-chi/jwtauth"
)

type GetChallenges struct {
	ChallengeRepository *repository.Challenge
	TokenAuth           *jwtauth.JWTAuth
}

func (u *GetChallenges) Execute(token string, ctx context.Context) ([]model.Challenge, error) {
	tokenAuth, err := model.DecodeToken(u.TokenAuth, ctx)
	if err != nil {
		return []model.Challenge{}, err
	}
	challenges, err := u.ChallengeRepository.FindByPerson(tokenAuth.ID)
	if err != nil {
		return []model.Challenge{}, err
	}
	return challenges, nil
}
