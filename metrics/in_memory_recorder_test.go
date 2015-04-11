package metrics_test

import (
	"time"

	"github.com/cghsystems/gosum/metrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("In Memory Metrics", func() {
	var (
		api   = "api"
		end   = time.Now()
		start = time.Now()
	)

	BeforeEach(func() {
		metrics.ResetInMemoryMetrics()
		metrics.RecordInMemory(api, start, end)
	})

	It("records exactly one result", func() {
		Ω(metrics.InMemoryMetrics).To(HaveLen(1))
	})

	It("records the expected api", func() {
		Ω(metrics.InMemoryMetrics[0].API).To(Equal(api))
	})

	It("records the expected start", func() {
		Ω(metrics.InMemoryMetrics[0].Start).To(Equal(start))
	})

	It("records the expected end", func() {
		Ω(metrics.InMemoryMetrics[0].End).To(Equal(end))
	})

})
