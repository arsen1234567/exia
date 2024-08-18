package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) clientErrorWithMessage(w http.ResponseWriter, message string, status int) {
	errorResponse := ErrorResponse{Error: message}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) sendResponseWithMessage(w http.ResponseWriter, message string, status int) {
	messageResponse := MessageResponse{Message: message}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(messageResponse)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
