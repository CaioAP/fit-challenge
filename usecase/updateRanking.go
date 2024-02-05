package usecase

import (
	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type UpdateRanking struct {
	RankingRepository *repository.Ranking
}

func (u *UpdateRanking) Execute(personId int, challengeId int) (model.Ranking, error) {
	ranking, err := u.RankingRepository.FindOne(challengeId, personId)
	if err != nil {
		return model.Ranking{}, err
	}
	ranking.Completed += 1
	return u.RankingRepository.Update(ranking)
}
