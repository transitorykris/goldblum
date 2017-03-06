package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server represents our API server
type Server struct {
	db  *sqlx.DB
	log *logrus.Logger
}

// NewServer creates a new Goldblum server
func NewServer() (*Server, error) {
	server := Server{
		log: logrus.New(),
	}
	return &server, nil
}

// ConnectDB connects our server to the given DB
func (s *Server) ConnectDB(db string) error {
	var err error
	s.db, err = sqlx.Connect("mysql", db)
	return err
}

// Router returns and HTTP router with the handlers for this server
func (s *Server) Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/editor", s.EditorHandler()).Methods("GET")
	r.Handle("/", s.DynamicHandler()).Methods("GET")
	return r
}

// Close closes down the server
func (s *Server) Close() error {
	return s.db.Close()
}
