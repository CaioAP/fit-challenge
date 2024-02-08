package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/caioap/desafio_bonde/dto"
	"github.com/caioap/desafio_bonde/usecase"
)

type Auth struct {
	LoginUsecase    *usecase.Login
	RegisterUsecase *usecase.Register
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var loginDto dto.Login
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.LoginUsecase.Execute(loginDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// Secure: true,
		Name:  "jwt",
		Value: output.Token,
	})
	response, err := json.Marshal(map[string]interface{}{
		"id":      output.ID,
		"message": "logged in",
		"status":  "ok",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (h *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		// Secure: true,
		Name:  "jwt",
		Value: "",
	})
	response, err := json.Marshal(map[string]interface{}{
		"message": "logged out",
		"status":  "ok",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, errors.New("method not allowed").Error(), http.StatusMethodNotAllowed)
		return
	}
	var registerDto dto.Register
	err := json.NewDecoder(r.Body).Decode(&registerDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.RegisterUsecase.Execute(registerDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// Secure: true,
		Name:  "jwt",
		Value: output.Token,
	})
	response, err := json.Marshal(map[string]interface{}{
		"id":      output.ID,
		"message": "registered",
		"status":  "ok",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
