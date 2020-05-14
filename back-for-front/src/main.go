package main

import (
	"net/http"

	"github.com/fsantiag/track-progress/back-for-front/src/configuration"
	"github.com/fsantiag/track-progress/back-for-front/src/server"
	"github.com/fsantiag/track-progress/back-for-front/src/util"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = &configuration.Logger
}

func main() {
	router := server.InitRouter()
	env := util.Getenv("PROFILE_ENV", "dev")

	logger.Debug("Environment: ", env)
	logger.Info("Server started....")
	logger.Fatal(http.ListenAndServe(util.Address[env], router))
}
