package main

import (
	"logger-service/data"
	"net/http"
)

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var logPayload LogPayload

	err := app.readJSON(w, r, &logPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// getting ready entry
	event := data.LogEntry{
		Name: logPayload.Name,
		Data: logPayload.Data,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
