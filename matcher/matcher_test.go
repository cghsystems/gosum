package matcher_test

import (
	. "github.com/gosum/matcher"
	. "github.com/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matcher", func() {

	var matcher Matcher

	Context(".DebitAmount", func() {
		BeforeEach(func() {
			matcher = DebitAmount()
		})

		It("matches a valid debit record with a positive debit amount", func() {
			record := Record{TransactionType: DEBIT, DebitAmount: float64(100)}
			Ω(matcher.IsExpectedTransactionType(record)).Should(BeTrue())
		})

		It("does not match a valid debit record with a zero debit amount", func() {
			record := Record{TransactionType: DEBIT, DebitAmount: float64(0)}
			Ω(matcher.IsExpectedTransactionType(record)).Should(BeFalse())
		})

		It("does not match a valid debit record with a negative debit amount", func() {
			record := Record{TransactionType: DEBIT, DebitAmount: float64(-100)}
			Ω(matcher.IsExpectedTransactionType(record)).Should(BeFalse())
		})

		It("returns the expected debit amount", func() {
			record := Record{TransactionType: DEBIT, DebitAmount: float64(100)}
			Ω(matcher.GetValue(record)).Should(Equal(float64(100)))
		})
	})

	Context(".CreditAmount", func() {
		BeforeEach(func() {
			matcher = CreditAmount()
		})

		It("returns the expected credit amount for a Credit transcation", func() {
			record := Record{TransactionType: CREDIT, CreditAmount: float64(10000)}
			Ω(matcher.GetValue(record)).Should(Equal(float64(10000)))
		})

		It("returns the expected credit amount for a FPI transaction", func() {
			record := Record{TransactionType: FPI, CreditAmount: float64(10000)}
			Ω(matcher.GetValue(record)).Should(Equal(float64(10000)))
		})
	})
})
