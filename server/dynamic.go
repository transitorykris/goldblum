package main

import (
	"fmt"
	"net/http"
	"plugin"

	gb "github.com/transitorykris/goldblum"
)

func (s *Server) lookupSO(method string, path string) (string, error) {
	var e Endpoint
	err := s.db.Get(&e, "SELECT `id`, `version` FROM endpoint WHERE `method`=? AND `path`=?", method, path)
	so := fmt.Sprintf("/%d/%d-%d.so", e.ID, e.ID, e.Version)
	s.log.Infoln("SO is at", so)
	return so, err
}

// DynamicHandler is a simple interface for modifying endpoints
func (s *Server) DynamicHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)
		so, err := s.lookupSO(r.Method, r.URL.Path)
		if err != nil {
			s.log.Errorln(err)
			gb.Response(w, &gb.EmptyResponse{}, http.StatusNotFound)
			return
		}
		p, err := plugin.Open(so)
		if err != nil {
			s.log.Errorln(err)
			gb.Response(w, &gb.EmptyResponse{}, http.StatusInternalServerError)
			return
		}
		handler, err := p.Lookup("Handler")
		if err != nil {
			s.log.Errorln(err)
			gb.Response(w, &gb.EmptyResponse{}, http.StatusInternalServerError)
			return
		}
		handler.(func(http.ResponseWriter, *http.Request))(w, r)
	})
}
