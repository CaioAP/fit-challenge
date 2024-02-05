package usecase

import (
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type GetRanking struct {
	ChallengeRepository repository.Challenge
	RankingRepository   repository.Ranking
}

func (u *GetRanking) Execute(challengeId int, personId int) (model.Ranking, error) {
	challenge, err := u.ChallengeRepository.FindById(challengeId)
	if err != nil {
		return model.Ranking{}, nil
	}
	rankings, err := u.RankingRepository.FindByChallenge(challengeId)
	if err != nil {
		return model.Ranking{}, err
	}
	var ranking model.Ranking
	for i, rank := range rankings {
		rank.Position = i + 1
		rank.Remaining = challenge.Goal - rank.Completed
		if rank.Person.ID == personId {
			ranking = rank
		}
	}
	return ranking, nil
}
