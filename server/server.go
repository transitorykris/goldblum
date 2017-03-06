package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
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
	if err != nil {
		return err
	}
	s.log.Println("Running migrations")
	s.migrations(db)
	return err
}

// Router returns and HTTP router with the handlers for this server
func (s *Server) Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/editor", s.EditorHandler()).Methods("GET")
	r.Handle("/editor/endpoint", s.CreateEndpointHandler()).Methods("POST")
	r.Handle("/editor/endpoint/{id:[0-9]+}", s.GetEndpointHandler()).Methods("GET")
	r.Handle("/editor/endpoint/{id:[0-9]+}", s.UpdateEndpointHandler()).Methods("POST")
	r.PathPrefix("/").Handler(s.DynamicHandler())
	return r
}

// Close closes down the server
func (s *Server) Close() error {
	return s.db.Close()
}

// migrations runs our database migrations
func (s *Server) migrations(openStr string) error {
	// Setup the goose configuration
	c := &goose.DBConf{
		MigrationsDir: "db/migrations",
		Env:           "development",
		Driver: goose.DBDriver{
			Name:    "mysql",
			OpenStr: openStr,
			Import:  "github.com/go-sql-driver/mysql",
			Dialect: &goose.MySqlDialect{},
		},
	}

	// Get the latest possible migration
	latest, err := goose.GetMostRecentDBVersion(c.MigrationsDir)
	if err != nil {
		return err
	}

	// Migrate up to the latest version
	if err = goose.RunMigrationsOnDb(c, c.MigrationsDir, latest, s.db.DB); err != nil {
		return err
	}

	return nil
}
