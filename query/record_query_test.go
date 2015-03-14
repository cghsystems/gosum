package query_test

import (
	"time"

	"github.com/cghsystems/gosum/data"
	. "github.com/cghsystems/gosum/matcher"
	. "github.com/cghsystems/gosum/query"
	"github.com/cghsystems/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordQuery", func() {
	const (
		Zero = float64(0)
	)

	var (
		records     record.Records
		recordQuery RecordQuery
		err         error
	)

	BeforeEach(func() {
		repo := data.NewFileRepository("assets/test_data.json")
		records, err = repo.LoadRecords()
		Ω(err).ShouldNot(HaveOccurred())
		recordQuery = NewRecordQuery(records)
	})

	Context(".SumDebitTotal", func() {
		It("calculates the Debit Amount total", func() {
			debitTotal := recordQuery.Sum(DebitAmount())
			Ω(debitTotal).Should(Equal(683582.8999999973))
		})
	})

	Context(".SumCreditTotal", func() {
		It("calculates the Credit Total", func() {
			creditTotal := recordQuery.Sum(CreditAmount())
			Ω(creditTotal).Should(Equal(1.6031996199999999e+06))
		})
	})

	Context(".DebitAmount", func() {
		var highestRecords record.Records

		BeforeEach(func() {
			highestRecords = recordQuery.FindHighest(DebitAmount())
		})

		Context(".FindHighest", func() {
			Context("finds multiple highest values", func() {
				It("should find two records", func() {
					records = record.Records{
						record.Record{
							DebitAmount:     1000,
							TransactionType: record.DEBIT,
						},
						record.Record{
							DebitAmount:     1000,
							TransactionType: record.DEBIT,
						},
					}

					recordQuery = NewRecordQuery(records)
					highestRecords = recordQuery.FindHighest(DebitAmount())
					Ω(highestRecords).Should(HaveLen(2))
				})
			})

			It("find a single record", func() {
				Ω(highestRecords).Should(HaveLen(1))
			})

			It("returns no error finding the first record", func() {
				_, err := highestRecords.First()
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("finds the highest debit amount", func() {
				highestRecord, _ := highestRecords.First()
				Ω(highestRecord.DebitAmount).Should(Equal(float64(28000)))
			})
		})

		Context(".FindLowest", func() {
			var lowestRecords record.Records

			BeforeEach(func() {
				lowestRecords = recordQuery.FindLowest(DebitAmount())
			})

			Context("finds multiple low values", func() {
				It("should find two records", func() {
					records = record.Records{
						record.Record{
							DebitAmount:     100,
							TransactionType: record.DEBIT,
						},
						record.Record{
							DebitAmount:     100,
							TransactionType: record.DEBIT,
						},
					}

					recordQuery = NewRecordQuery(records)
					lowestRecords = recordQuery.FindLowest(DebitAmount())
					Ω(lowestRecords).Should(HaveLen(2))
				})
			})

			It("finds 1 records", func() {
				Ω(lowestRecords).Should(HaveLen(1))
			})

			It("lowest records have debit value of 0.12", func() {
				for _, lowestRecord := range lowestRecords {
					Ω(lowestRecord.DebitAmount).Should(Equal(float64(0.13)))
				}
			})

			It("finds the lowest debit amount", func() {
				lowestRecord, _ := lowestRecords.First()
				Ω(lowestRecord.DebitAmount).Should(Equal(0.13))
			})

			It("does not error finding the first record", func() {
				_, err := lowestRecords.First()
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Context(".RecordsWithDescription", func() {
		It("finds the records matching the description", func() {
			records := recordQuery.RecordsWithDescription("TESCO").Records()
			Ω(records).Should(HaveLen(98))
		})

		It("ignores the case of the description text", func() {
			records := recordQuery.RecordsWithDescription("tesco").Records()
			Ω(records).Should(HaveLen(98))
		})

		It("returns zero results if the description text cannot be found", func() {
			records := recordQuery.RecordsWithDescription("wibble").Records()
			Ω(records).Should(BeEmpty())
		})
	})

	Context(".RecordsInDataRange", func() {
		BeforeEach(func() {
			startDate, _ := time.Parse(time.RFC3339, "2009-01-01T00:00:00+00:00")
			endDate, _ := time.Parse(time.RFC3339, "2009-01-31T00:00:00+00:00")
			recordQuery = recordQuery.RecordsInDateRange(startDate, endDate)
		})

		It("finds all of the records from January 2009", func() {
			Ω(61).To(Equal(len(recordQuery.Records())))
		})

		It("sums the debit amount in January 2009", func() {
			Ω(3847.36).To(Equal(recordQuery.Sum(DebitAmount())))
		})

		Context(".DebitAmountMatcher", func() {
			Context("the highest debit total in January 2009", func() {
				var highestRecords record.Records

				BeforeEach(func() {
					startDate, _ := time.Parse(time.RFC3339, "2009-01-01T00:00:00+00:00")
					endDate, _ := time.Parse(time.RFC3339, "2009-01-31T00:00:00+00:00")
					recordQuery = recordQuery.RecordsInDateRange(startDate, endDate)
					highestRecords = recordQuery.FindHighest(DebitAmount())
				})

				It("contains only one record", func() {
					Ω(highestRecords).Should(HaveLen(1))
				})

				It("the record has a value of 2329", func() {
					highestRecord := highestRecords[0]
					debitAmount := highestRecord.DebitAmount
					Ω(float64(891.9)).Should(Equal(debitAmount))
				})
			})
		})

		Context(".Mean", func() {
			It("calculates the mean DebitAmount", func() {
				meanDebitAmount := recordQuery.Mean(DebitAmount())
				Ω(meanDebitAmount).Should(Equal(float64(192.368)))
			})

			It("calculates the mean CreditAmount", func() {
				meanDebitAmount := recordQuery.Mean(CreditAmount())
				Ω(meanDebitAmount).Should(Equal(516.75))
			})

			Context("no records with a credit transaction type", func() {
				BeforeEach(func() {
					recordQuery = NewRecordQuery(record.Records{
						record.Record{
							TransactionType: "DEB",
						},
					})
				})

				It("has a zero Mean value", func() {
					meanDebitAmount := recordQuery.Mean(CreditAmount())
					Ω(meanDebitAmount).Should(BeZero())
				})
			})
		})

		Context(".CreditAmountMatcher", func() {
			Context("the highest credit total in February 2009", func() {
				var highestRecords record.Records

				BeforeEach(func() {
					highestRecords = recordQuery.FindHighest(CreditAmount())
				})

				It("contains only one record", func() {
					Ω(highestRecords).Should(HaveLen(1))
				})

				It("finds the highest credit transaction types in February 2009", func() {
					highestRecord, _ := highestRecords.First()
					Ω(float64(1715)).To(Equal(highestRecord.CreditAmount))
				})
			})
		})
	})

	Context("Chaining searches", func() {
		It("finds the total Debit Amount spent in Tesco", func() {
			totalSpentInTesco := recordQuery.
				RecordsWithDescription("TESCO").
				Sum(DebitAmount())
			Ω(float64(6332.4400000000005)).Should(Equal(totalSpentInTesco))
		})

		It("finds the record with highest Debit Amount spent in Tesco", func() {
			record, _ := recordQuery.
				RecordsWithDescription("TESCO").
				FindHighest(DebitAmount()).First()
			Ω(float64(678.4)).Should(Equal(record.DebitAmount))
		})

		It("finds the record with lowest Debit Amount spent in Tesco", func() {
			record, _ := recordQuery.
				RecordsWithDescription("TESCO").
				FindLowest(DebitAmount()).First()

			Ω(float64(3.92)).Should(Equal(record.DebitAmount))
		})
	})

	Context("No results are returned from search", func() {
		var (
			err    error
			record record.Record
		)

		BeforeEach(func() {
			record, err = recordQuery.
				RecordsWithDescription("Wibble").
				FindLowest(DebitAmount()).First()
		})

		It("returns NilRecord", func() {
			Ω(record.IsNilRecord()).Should(BeTrue())
		})

		It("returns error", func() {
			Ω(err).Should(HaveOccurred())
		})
	})
})
