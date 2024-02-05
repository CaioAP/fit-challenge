package usecase

import (
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type GetRankings struct {
	ChallengeRepository repository.Challenge
	RankingRepository   repository.Ranking
}

func (u *GetRankings) Execute(challengeId int, personId int) ([]model.Ranking, error) {
	challenge, err := u.ChallengeRepository.FindById(challengeId)
	if err != nil {
		return []model.Ranking{}, nil
	}
	rankings, err := u.RankingRepository.FindByChallenge(challengeId)
	if err != nil {
		return []model.Ranking{}, err
	}
	for i, rank := range rankings {
		rank.Position = i + 1
		rank.Remaining = challenge.Goal - rank.Completed
	}
	return rankings, nil
}
