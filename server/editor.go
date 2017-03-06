package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	gb "github.com/transitorykris/goldblum"
)

// Endpoint represents a single endpoint created by the user
type Endpoint struct {
	ID      int64
	Version int64
	Method  string
	Path    string
	Code    string
}

func (s *Server) getEndpoints() ([]Endpoint, error) {
	var e []Endpoint
	err := s.db.Select(&e, "SELECT `id`, `version`, `method`, `path`, `code` FROM `endpoint`")
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

func (s *Server) createEndpoint(method string, path string, code string) (Endpoint, error) {
	var e Endpoint
	res, err := s.db.Exec(
		"INSERT INTO `endpoint` (`method`, `path`, `code`) VALUES (?, ?, ?)",
		method, path, code,
	)
	if err != nil {
		return e, err
	}
	id, _ := res.LastInsertId()
	e = Endpoint{
		ID:     id,
		Method: method,
		Path:   path,
		Code:   code,
	}
	s.log.Errorln("WTF CODE?", e.Code)
	return e, nil
}

// CreateEndpointHandler is a simple interface for modifying endpoints
func (s *Server) CreateEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		r.ParseForm()
		endpoint, err := s.createEndpoint(r.FormValue("method"), r.FormValue("path"), r.FormValue("code"))
		if err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		s.log.Infoln("Compiling", endpoint)
		_, err = s.compile(endpoint.ID)
		if err != nil {
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/editor", 302)
	})
}

func (s *Server) getEndpoint(id int64) (Endpoint, error) {
	var e Endpoint
	err := s.db.Get(&e, "SELECT `id`, `version`, `method`, `path`, `code` FROM `endpoint` WHERE `id`=?", id)
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
	var version int64
	if err := s.db.Get(&version, "SELECT `version` FROM `endpoint` WHERE `id`=?", id); err != nil {
		return err
	}
	_, err := s.db.Exec("UPDATE `endpoint` SET `code`=?, `version`=? WHERE `id`=?", code, version+1, id)
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
			s.log.Errorln(err)
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		_, err := s.compile(id)
		if err != nil {
			s.log.Errorln(err)
			gb.Response(w, &gb.ErrorResponse{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/editor", 302)
	})
}
