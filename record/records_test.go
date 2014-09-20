package record_test

import (
	"github.com/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Records", func() {
	var records record.Records

	Context("There are zero records", func() {
		BeforeEach(func() {
			records = record.Records{}
		})

		Context(".First", func() {
			It("returns an error", func() {
				_, err := records.First()
				Ω(err).Should(HaveOccurred())
			})
		})

		Context(".Delete", func() {
			It("returns an error", func() {
				toDelete := testRecord("To Delete")
				_, err := records.Delete(toDelete)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Context("There are many records", func() {
		var firstRecord, secondRecord, thirdRecord record.Record

		BeforeEach(func() {
			firstRecord = testRecord("First Record")
			secondRecord = testRecord("Second Record")
			thirdRecord = testRecord("Third Record")
			records = record.Records{
				firstRecord,
				secondRecord,
				thirdRecord,
			}
		})

		Context(".FirstRecord", func() {
			It("will return the first record in the slice", func() {
				Ω(records.First()).Should(Equal(firstRecord))
			})
		})

		Context(".DeleteAll", func() {
			It("deletes all data", func() {
				Ω(records.DeleteAll()).Should(BeEmpty())
			})
		})

		Context(".Delete", func() {
			var err error

			BeforeEach(func() {
				records, err = records.Delete(firstRecord)
			})

			Context("Successful delete", func() {
				It("will not error", func() {
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("will not contain the deleted record", func() {
					Ω(records).ShouldNot(ContainElement(firstRecord))
				})

				It("will have the expected length", func() {
					Ω(records).Should(HaveLen(2))
				})

				It("will contains the expcted records", func() {
					Ω(records).Should(ConsistOf(secondRecord, thirdRecord))
				})
			})

			Context("unexpected record", func() {
				It("will return an error if there is no record to delete", func() {
					_, err := records.Delete(testRecord("record does not exist"))
					Ω(err).Should(HaveOccurred())
				})
			})
		})
	})
})

func testRecord(transactionDescription string) record.Record {
	record := record.NilRecord()
	record.TransactionDescription = transactionDescription
	return record
}
