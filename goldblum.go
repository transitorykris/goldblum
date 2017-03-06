package goldblum

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

// Goldblum is the the struct that gets passed to each handler with some goodies
type Goldblum struct {
	DB  *sqlx.DB
	Log *logrus.Logger
}

// Handler is an HTTP handler created by the user
type Handler func(*Goldblum, http.ResponseWriter, *http.Request)

// ErrorResponse is used when a json object needs to be returned with just an error
type ErrorResponse struct {
	Error string `json:"error"`
}

// EmptyResponse is used when we need no body
type EmptyResponse struct{}

// Response is a nice wrapper for sending JSON responses
func Response(w http.ResponseWriter, v interface{}, status int) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	fmt.Fprint(w, string(body))
	return nil
}
