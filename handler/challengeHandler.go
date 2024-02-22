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
	GetChallenges   usecase.GetChallenges
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

func (h Challenge) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
	}
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	challenges, err := h.GetChallenges.Execute(cookie.Value, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	response, err := json.Marshal(challenges)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(response)
}
