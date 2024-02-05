package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/caioap/desafio_bonde/handler"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/caioap/desafio_bonde/usecase"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5433 user=test password=testingdb dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("Unable to connect to database")
	}
	defer db.Close()

	personRepository := repository.Person{DB: db}
	challengeRepository := repository.Challenge{DB: db}
	activityRepository := repository.Activity{DB: db}
	rankingRepository := repository.Ranking{DB: db}

	createPersonUsecase := usecase.CreatePerson{
		PersonRepository: &personRepository,
	}
	getPersonUsecase := usecase.GetPerson{
		PersonRepository: &personRepository,
	}
	createRanking := usecase.CreateRanking{
		RankingRepository: &rankingRepository,
	}
	updateRanking := usecase.UpdateRanking{
		RankingRepository: &rankingRepository,
	}
	createChallengeUsecase := usecase.CreateChallenge{
		PersonRepository:    &personRepository,
		ChallengeRepository: &challengeRepository,
		CreateRanking:       &createRanking,
	}
	createActivity := usecase.CreateActivity{
		PersonRepository:    &personRepository,
		ChallengeRepository: &challengeRepository,
		ActivityRepository:  &activityRepository,
		UpdateRanking:       &updateRanking,
	}

	r := chi.NewRouter()

	personHandler := handler.Person{
		CreatePerson: &createPersonUsecase,
		GetPerson:    &getPersonUsecase,
	}
	r.Route("/people", func(r chi.Router) {
		r.Post("/", personHandler.Create)
		r.Get("/{id}", personHandler.Get)
	})

	challengeHandler := handler.Challenge{CreateChallenge: createChallengeUsecase}
	r.Route("/challenges", func(r chi.Router) {
		r.Post("/", challengeHandler.Create)
	})

	activityHandler := handler.Activity{CreateActivity: createActivity}
	r.Route("/activities", func(r chi.Router) {
		r.Post("/", activityHandler.Create)
	})

	log.Println("server running at port 9000...")
	http.ListenAndServe(":9000", r)
}
