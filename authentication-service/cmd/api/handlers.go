package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.Models.User.GetByEmail(req.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credientails"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(req.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credientails"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
