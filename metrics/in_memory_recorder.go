package metrics

import "time"

type metric struct {
	API   string
	Start time.Time
	End   time.Time
}

var InMemoryMetrics = []metric{}

func RecordInMemory(api string, start, end time.Time) {
	InMemoryMetrics = append(InMemoryMetrics, metric{api, start, end})
}

func ResetInMemoryMetrics() {
	InMemoryMetrics = []metric{}
}
