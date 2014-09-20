package record_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gosum/record"
)

var _ = Describe("Record", func() {

	Context("NilRecord", func() {
		It("evaluates NilRecord to true", func() {
			record := NilRecord()
			Ω(record.IsNilRecord()).Should(BeTrue())
		})

		It("evaluates empty record to true", func() {
			record := Record{}
			Ω(record.IsNilRecord()).Should(BeTrue())
		})

		It("evaluates a record with data to false", func() {
			record := Record{TransactionDescription: "Not a NilRecord"}
			Ω(record.IsNilRecord()).Should(BeFalse())
		})
	})
})
