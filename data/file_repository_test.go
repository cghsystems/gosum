package data_test

import (
	"github.com/cghsystems/gosum/data"
	"github.com/cghsystems/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filedata", func() {

	Context(".LoadRecords", func() {
		var err error
		var records record.Records
		var dataRepository data.Repository

		BeforeEach(func() {
			dataRepository = data.NewFileRepository("../query/assets/test_data.json")
		})

		Context("invalid reposotry file provided", func() {
			It("should return an error", func() {
				dataRepository = data.NewFileRepository("/path/to/nothing")
				_, err = dataRepository.LoadRecords()
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("valid data file provided", func() {
			BeforeEach(func() {
				records, err = dataRepository.LoadRecords()
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
