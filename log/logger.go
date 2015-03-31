package log

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

func Init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	logging.AddModuleLevel(backend)
	logging.SetBackend(backend)
}

func Info(message string) {
	log.Info(message)
}

func Debug(message string) {
	log.Debug(message)
}
