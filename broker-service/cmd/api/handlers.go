package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type (
	RequestPayload struct {
		Action string      `json:"action"`
		Auth   AuthPayload `json:"auth"`
	}

	AuthPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload

	err := app.readJSON(w, r, req)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch req.Action {
	case "auth":
		app.authenticate(w, req.Auth)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// check statius code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service "))
		return
	}

	// json from service
	data := JsonResponse{}

	// decode json from service
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if data.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	payload := JsonResponse{
		Error:   data.Error,
		Message: "Authenticated",
		Data:    data.Data,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
