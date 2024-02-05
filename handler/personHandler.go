package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/usecase"
	"github.com/go-chi/chi/v5"
)

type Person struct {
	CreatePerson *usecase.CreatePerson
	GetPerson    *usecase.GetPerson
}

func (h *Person) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var personDto dto.PersonCreate
	err := json.NewDecoder(r.Body).Decode(&personDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	person, err := h.CreatePerson.Execute(personDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (h *Person) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, errors.New("invalid id").Error(), http.StatusBadRequest)
		return
	}
	person, err := h.GetPerson.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	response, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
