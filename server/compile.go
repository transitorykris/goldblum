package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func (s *Server) compile(id int64) (string, error) {
	endpoint, err := s.getEndpoint(id)
	if err != nil {
		s.log.Errorln(err)
		return "", err
	}
	dir := fmt.Sprintf("/%d", id)
	os.Mkdir(dir, 0644)
	if err = ioutil.WriteFile(fmt.Sprintf("%s/%d.go", dir, id), []byte(endpoint.Code), 0644); err != nil {
		s.log.Errorln(err)
		return "", err
	}
	so := fmt.Sprintf("%s/%d.so", dir, id)
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", so)
	cmd.Dir = dir
	if err = cmd.Run(); err != nil {
		s.log.Errorln(err)
		return "", err
	}
	return so, err
}
