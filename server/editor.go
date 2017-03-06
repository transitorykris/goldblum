package main

import (
	"net/http"

	gb "github.com/transitorykris/goldblum"
)

// EditorHandler is a simple interface for modifying endpoints
func (s *Server) EditorHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}

// CreateEndpointHandler is a simple interface for modifying endpoints
func (s *Server) CreateEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}

// GetEndpointHandler is a simple interface for modifying endpoints
func (s *Server) GetEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}

// UpdateEndpointHandler is a simple interface for modifying endpoints
func (s *Server) UpdateEndpointHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}
