package log

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")
var backend = logging.NewLogBackend(os.Stderr, "", 0)

func Info(message string) {
	logging.AddModuleLevel(backend)
	logging.SetBackend(backend)
	log.Info(message)
}

func Debug(message string) {
	log.Debug(message)
}
