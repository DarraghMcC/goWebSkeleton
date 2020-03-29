package main

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func getPort() string {
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8090"
	}

	return ":" + port
}

func isTimeStampDisabled() bool {
	timeStamp := os.Getenv("DISABLE_LOG_TIMESTAMP")

	if timeStamp == "" {
		return false
	}

	v, err := strconv.ParseBool(timeStamp)
	if err != nil {
		log.Error("Invalid `DISABLE_LOG_TIMESTAMP` set, this should be a boolean", err.Error())
		return false
	}

	return v
}
