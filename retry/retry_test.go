package retry_test

import (
	"github.com/cghsystems/gosum/retry"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Retry", func() {
	Context("successful retries", func() {
		var (
			actualRetries = 0
			err           error
		)

		BeforeEach(func() {
			err = retry.Execute(10, func() {
				actualRetries++
			})
		})

		It("retries 10 times", func() {
			Ω(actualRetries).Should(Equal(10))
		})

		It("does not return error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
