package metrics

import (
	"fmt"
	"os"
	"time"

	"github.com/cghsystems/gosum/log"
)

var db *os.File

func InitFileRecorder(file string) {
	db, _ = os.Create(file)
}

func Record(api string, requestTime time.Time) {
	log.Info("Writing to file" + db.Name())
	_, err := db.WriteString(fmt.Sprintf("%v,%v", api, requestTime))
	if err != nil {
		log.Info(err.Error())
	}
	db.Close()
}
