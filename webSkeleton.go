//Basic web app skeleton
package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var GitCommit string
var BuildDate string

func main() {
	setUp()

	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())

	port := getPort()

	log.WithFields(log.Fields{"gitCommit": GitCommit, "buildDate": BuildDate, "port": port}).Info("Web Skeleton started")

	_ = http.ListenAndServe(port, nil)
}

func setUp() {
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: isTimeStampDisabled(),
	})

	if err := godotenv.Load(); err != nil {
		log.Warn(err.Error())
	}

	initialiseMetrics()
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("pong"))

	if err != nil {
		log.Error("Ping response failure:", err.Error())
	}
}
