package main

import (
	"fmt"
	"net/http"
	"plugin"

	gb "github.com/transitorykris/goldblum"
)

func (s *Server) lookupEndpoint(method string, path string) (int64, error) {
	var id int64
	err := s.db.Get(&id, "SELECT `id` FROM endpoint WHERE `method`=? AND `path`=?", method, path)
	return id, err
}

// DynamicHandler is a simple interface for modifying endpoints
func (s *Server) DynamicHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Debugln(r.Method, r.URL.Path, r.RemoteAddr)

		id, err := s.lookupEndpoint(r.Method, r.URL.Path)
		if err != nil {
			gb.Response(w, &gb.EmptyResponse{}, http.StatusNotFound)
			return
		}
		s.log.Infoln("Would have executed endpoint", id)

		p, err := plugin.Open(fmt.Sprintf("/%d/%d.so", id, id))
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
