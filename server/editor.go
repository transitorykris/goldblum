package main

import (
	"net/http"
	"strconv"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	gb "github.com/transitorykris/goldblum"
)

// Endpoint represents a single endpoint created by the user
type Endpoint struct {
	ID     int64
	Method string
	Path   string
	Code   string
}

func (s *Server) getEndpoints() ([]Endpoint, error) {
	var e []Endpoint
	err := s.db.Select(&e, "SELECT `id`, `method`, `path`, `code` FROM `endpoint`")
	return e, err
}

// EditorHandler is a simple interface for modifying endpoints
func (s *Server) EditorHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		endpoints, err := s.getEndpoints()
		if err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		t, err := template.ParseFiles("template/endpoints.html")
		if err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		t.Execute(w, endpoints)
	})
}

// CreateEndpointHandler is a simple interface for modifying endpoints
func (s *Server) CreateEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}

func (s *Server) getEndpoint(id int64) (Endpoint, error) {
	var e Endpoint
	err := s.db.Get(&e, "SELECT `id`, `method`, `path`, `code` FROM `endpoint`")
	return e, err
}

// GetEndpointHandler is a simple interface for modifying endpoints
func (s *Server) GetEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		endpoint, err := s.getEndpoint(id)
		t, err := template.ParseFiles("template/endpoint.html")
		if err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		t.Execute(w, endpoint)
	})
}

func (s *Server) updateEndpoint(id int64, code string) error {
	_, err := s.db.Exec("UPDATE `endpoint` SET `code`=? WHERE `id`=?", code, id)
	return err
}

// UpdateEndpointHandler is a simple interface for modifying endpoints
func (s *Server) UpdateEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		r.ParseForm()
		code := r.FormValue("code")
		if err := s.updateEndpoint(id, code); err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/editor", 302)
	})
}
