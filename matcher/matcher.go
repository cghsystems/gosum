package matcher

import "github.com/cghsystems/gosum/record"

type Matcher interface {
	GetValue(record record.Record) float64
	IsExpectedTransactionType(record record.Record) bool
}

//Credit Amount
type creditAmountMatcher struct{}

func CreditAmount() Matcher {
	return creditAmountMatcher{}
}

func (creditAmountMatcher creditAmountMatcher) IsExpectedTransactionType(record record.Record) bool {
	return record.IsCredit() || record.IsFPI()
}

func (creditAmountMatcher creditAmountMatcher) GetValue(record record.Record) float64 {
	return record.CreditAmount
}

//Debit Amount
type debitAmountMatcher struct{}

func DebitAmount() Matcher {
	return debitAmountMatcher{}
}

func (debitAmountMatcher debitAmountMatcher) IsExpectedTransactionType(record record.Record) bool {
	return record.IsDebit() && record.DebitAmount > 0
}

func (debitAmountMatcher debitAmountMatcher) GetValue(record record.Record) float64 {
	return record.DebitAmount
}
