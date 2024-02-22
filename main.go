package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/caioap/desafio_bonde/handler"
	"github.com/caioap/desafio_bonde/repository"
	"github.com/caioap/desafio_bonde/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	connStr := "host=localhost port=5433 user=test password=testingdb dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("Unable to connect to database")
	}
	defer db.Close()

	const JWTSecret = "fit-3286-challenge-1389"
	tokenAuth := jwtauth.New("HS256", []byte(JWTSecret), nil)

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
	getPersonAuthorizedUsecase := usecase.GetPersonAuthorized{
		TokenAuth:        tokenAuth,
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
	getChallengesUsecase := usecase.GetChallenges{
		ChallengeRepository: &challengeRepository,
		TokenAuth:           tokenAuth,
	}
	createActivity := usecase.CreateActivity{
		PersonRepository:    &personRepository,
		ChallengeRepository: &challengeRepository,
		ActivityRepository:  &activityRepository,
		UpdateRanking:       &updateRanking,
	}
	loginUsecase := usecase.Login{
		TokenAuth:        tokenAuth,
		PersonRepository: &personRepository,
	}
	registerUsecase := usecase.Register{
		TokenAuth:        tokenAuth,
		PersonRepository: &personRepository,
	}

	authHandler := handler.Auth{
		LoginUsecase:    &loginUsecase,
		RegisterUsecase: &registerUsecase,
	}
	personHandler := handler.Person{
		CreatePerson:        &createPersonUsecase,
		GetPerson:           &getPersonUsecase,
		GetPersonAuthorized: &getPersonAuthorizedUsecase,
	}
	challengeHandler := handler.Challenge{CreateChallenge: createChallengeUsecase, GetChallenges: getChallengesUsecase}
	activityHandler := handler.Activity{CreateActivity: createActivity}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/me", personHandler.Me)

		r.Route("/people", func(r chi.Router) {
			r.Get("/{id}", personHandler.Get)
		})
		r.Route("/challenges", func(r chi.Router) {
			r.Get("/", challengeHandler.GetAll)
			r.Post("/", challengeHandler.Create)
		})
		r.Route("/activities", func(r chi.Router) {
			r.Post("/", activityHandler.Create)
		})
	})

	r.Group(func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/logout", authHandler.Logout)
		r.Post("/register", authHandler.Register)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	s := c.Handler(r)

	log.Println("server running at port 9000...")
	log.Fatal(http.ListenAndServe(":9000", s))
}
