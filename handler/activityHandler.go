package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/usecase"
)

type Activity struct {
	CreateActivity usecase.CreateActivity
}

func (h Activity) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var activityDto dto.ActivityCreate
	err := json.NewDecoder(r.Body).Decode(&activityDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	activity, err := h.CreateActivity.Execute(activityDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
