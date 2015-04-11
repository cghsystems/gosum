package metrics

import (
	"fmt"
	"os"
	"time"

	"github.com/cghsystems/gosum/log"
)

var (
	db  *os.File
	err error
)

func InitFileRecorder(file string) error {
	db, err = os.Create(file)
	return err
}

func RecordInFile(api string, startTime, endTime time.Time) {
	log.Info("Writing to file" + db.Name())
	_, err := db.WriteString(fmt.Sprintf("%v,%v,%v", api, startTime, endTime))
	if err != nil {
		log.Info(err.Error())
	}
	db.Close()
}
