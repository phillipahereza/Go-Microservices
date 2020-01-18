package service

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func StartWebServer(port string) {
	logrus.Infof("Starting HTTP service at %v", port)
	r := NewRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, nil) //Gorountine will block here
	if err != nil {
		logrus.Errorln("An error occured starting HTTP listener at port " + port)
		logrus.Errorln("Error: " + err.Error())
	}
}
