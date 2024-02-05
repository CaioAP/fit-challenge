package usecase

import (
	"log"

	"github.com/caioap/desafio_bonde/model"
	"github.com/caioap/desafio_bonde/repository"
)

type CreateRanking struct {
	RankingRepository *repository.Ranking
}

func (u *CreateRanking) Execute(person model.Person, challenge model.Challenge) (model.Ranking, error) {
	ranking, err := u.RankingRepository.FindOne(challenge.ID, person.ID)
	if err.Error() != "ranking not found" {
		return model.Ranking{}, err
	}
	if ranking.ID != 0 {
		return ranking, nil
	}
	ranking = model.Ranking{
		Completed: 0,
		Person:    person,
		Challenge: challenge,
	}
	id, err := u.RankingRepository.Create(ranking)
	if err != nil {
		log.Println(err)
		return model.Ranking{}, err
	}
	ranking.ID = id
	return ranking, nil
}
