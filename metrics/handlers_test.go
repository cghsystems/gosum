package metrics_test

import (
	"time"

	"github.com/cghsystems/gosum/metrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	BeforeEach(func() {
		metrics.ResetInMemoryMetrics()
	})

	It("records metrics in handlers", func() {
		metrics.RecordFunc(metrics.RecordInMemory)
		metrics.RecordFunc(metrics.RecordInMemory)
		metrics.Record("test", time.Now(), time.Now())
		Ω(metrics.InMemoryMetrics).To(HaveLen(2))
	})
})
