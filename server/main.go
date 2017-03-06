package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
)

type specification struct {
	Bind string `envconfig:"bind" default:":8001"`
	DB   string `envconfig:"db" default:"root:secret@tcp(mysql:3306)/goldblum"`
}

func main() {
	var err error
	logger := logrus.New()

	var spec specification
	if err = envconfig.Process("APP", &spec); err != nil {
		logger.Fatalln(err)
	}
	logger.Info(spec)

	s, err := NewServer()
	if err != nil {
		logger.Fatalln(err)
	}

	for {
		if err = s.ConnectDB(spec.DB); err != nil {
			logger.WithField("func", "main").Warnln("Problem connecting to DB", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	defer s.Close()

	logger.WithField("func", "main").Info("Starting")
	err = http.ListenAndServe(spec.Bind, s.Router())
	if err != nil {
		logger.Errorln(err)
	}
}
