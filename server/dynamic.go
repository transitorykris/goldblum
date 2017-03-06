package main

import (
	"net/http"

	gb "github.com/transitorykris/goldblum"
)

// DynamicHandler is a simple interface for modifying endpoints
func (s *Server) DynamicHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		gb.Response(w, &gb.EmptyResponse{}, http.StatusOK)
	})
}
