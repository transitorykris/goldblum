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
	src := fmt.Sprintf("%s/%d-%d.go", dir, id, endpoint.Version)
	err = ioutil.WriteFile(src, []byte(endpoint.Code), 0644)
	if err != nil {
		s.log.Errorln(err)
		return "", err
	}
	so := fmt.Sprintf("%s/%d-%d.so", dir, id, endpoint.Version)
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", so, src)
	if err = cmd.Run(); err != nil {
		s.log.Errorln(err)
		return "", err
	}
	return so, err
}
