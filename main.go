package main

import (
	"os"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/op/go-logging"
)

func initLogger() *logging.Logger {
	// Define the log format
	logFormat := logging.MustStringFormatter(
		"%{color}%{time:2006-01-02 15:04:05.000} ▶ %{level:.4s} %{message} (in %{shortfunc})",
	)

	// Create a log backend that writes to standard error
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Apply the log format to the backend
	backendFormatter := logging.NewBackendFormatter(backend, logFormat)

	// Set the backend as the logging backend
	logging.SetBackend(backendFormatter)

	// Enable Debug level logs
	logging.SetLevel(logging.DEBUG, "")

	// Create and return the logger
	return logging.MustGetLogger("docker-on-top")
}

var log *logging.Logger = initLogger()

func main() {
	// TODO: cleanup active mounts before starting

	log.Info("Hello there!")

	handler := volume.NewHandler(MustNewDockerOnTop("/var/lib/docker-on-top/"))
	log.Info(handler.ServeUnix("/run/docker/plugins/docker-on-top.sock", 0))

	// TODO: in case of abrupt termination, delete the socket file
}
