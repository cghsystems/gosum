package repository_test

import (
	"github.com/gosum/record"
	"github.com/gosum/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileRepository", func() {

	Context(".LoadRecords", func() {
		var err error
		var records record.Records

		Context("invalid reposotry file provided", func() {
			It("should return an error", func() {
				_, err = repository.LoadRecords("/path/to/invalid/file/")
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("valid repository file provided", func() {
			BeforeEach(func() {
				records, err = repository.LoadRecords("../query/assets/test_data.json")
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
