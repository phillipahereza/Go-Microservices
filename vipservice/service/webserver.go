package service

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	logrus.Infoln("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		logrus.Errorf("An error occured starting HTTP listener at port " + port)
		logrus.Errorf("Error: " + err.Error())
	}
}
