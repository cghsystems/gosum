package data_test

import (
	"github.com/gosum/data"
	"github.com/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filedata", func() {

	Context(".LoadRecords", func() {
		var err error
		var records record.Records

		Context("invalid reposotry file provided", func() {
			It("should return an error", func() {
				_, err = data.LoadRecords("/path/to/invalid/file/")
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("valid data file provided", func() {
			BeforeEach(func() {
				records, err = data.LoadRecords("../query/assets/test_data.json")
			})

			It("should load the expected number of records", func() {
				Ω(len(records)).Should(Equal(6241))
			})

			It("should load json file without error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
