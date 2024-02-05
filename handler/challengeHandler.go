package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/usecase"
)

type Challenge struct {
	CreateChallenge usecase.CreateChallenge
}

func (h Challenge) Show(w http.ResponseWriter, r *http.Request) {
	// body := page.IndexPage()
	// page := layout.IndexLayout("Go Templ + HTMX", body)

	// err := page.Render(r.Context(), w)
	// if err != nil {
	// 	log.Println("error", err)
	// }
}

func (h Challenge) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var challengeDto dto.ChallengeCreate
	err := json.NewDecoder(r.Body).Decode(&challengeDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	challenge, err := h.CreateChallenge.Execute(challengeDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(challenge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(response)
}
