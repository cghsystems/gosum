package metrics

import "time"

type record func(api string, start, end time.Time)

var recorders = []record{}

func RecordFunc(recordFunc record) {
	recorders = append(recorders, recordFunc)
}

func Record(api string, start, end time.Time) {
	for _, recorder := range recorders {
		recorder(api, start, end)
	}
}
