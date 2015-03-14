package query

import (
	"strings"
	"time"

	. "github.com/cghsystems/gosum/matcher"
	"github.com/cghsystems/gosum/record"
)

type RecordQuery interface {
	Records() record.Records
	Mean(Matcher) float64
	Sum(Matcher) float64
	FindLowest(Matcher) record.Records
	FindHighest(Matcher) record.Records
	RecordsInDateRange(startDate, endDate time.Time) RecordQuery
	RecordsWithDescription(searchString string) RecordQuery
}

type recordQuery struct {
	records record.Records
}

func NewRecordQuery(records record.Records) RecordQuery {
	return &recordQuery{records}
}

func (recordQuery *recordQuery) Records() record.Records {
	return recordQuery.records
}

func (recordQuery *recordQuery) Mean(matcher Matcher) float64 {
	count := float64(0)
	total := recordQuery.forEachRecord(func(record record.Record, total *float64) {
		if matcher.IsExpectedTransactionType(record) {
			count++
			*total += matcher.GetValue(record)
		}
	})

	if count > 0 {
		return total / count
	} else {
		return float64(0)
	}
}

func (recordQuery *recordQuery) Sum(matcher Matcher) float64 {
	return recordQuery.forEachRecord(func(record record.Record, total *float64) {
		if matcher.IsExpectedTransactionType(record) {
			*total += matcher.GetValue(record)
		}
	})
}

func (recordQuery *recordQuery) FindHighest(matcher Matcher) record.Records {

	winningRecord := record.Record{
		DebitAmount:  float64(-99999999),
		CreditAmount: float64(-99999999),
		Balance:      float64(-99999999),
	}

	results := record.Records{}

	for _, targetRecord := range recordQuery.records {
		if !matcher.IsExpectedTransactionType(targetRecord) {
			continue
		}

		if matcher.GetValue(targetRecord) > matcher.GetValue(winningRecord) {
			results = results.DeleteAll()
			results = append(results, targetRecord)
			winningRecord = targetRecord
		} else if matcher.GetValue(targetRecord) == matcher.GetValue(winningRecord) {
			results = append(results, targetRecord)
			winningRecord = targetRecord
		}
	}

	return results
}

func (recordQuery *recordQuery) FindLowest(matcher Matcher) record.Records {
	winningRecord := record.Record{
		DebitAmount:  float64(99999999),
		CreditAmount: float64(99999999),
		Balance:      float64(99999999),
	}

	results := record.Records{}

	for _, targetRecord := range recordQuery.records {
		if !matcher.IsExpectedTransactionType(targetRecord) {
			continue
		}

		if matcher.IsExpectedTransactionType(targetRecord) && matcher.GetValue(targetRecord) < matcher.GetValue(winningRecord) {
			results = results.DeleteAll()
			results = append(results, targetRecord)
			winningRecord = targetRecord
		} else if matcher.IsExpectedTransactionType(targetRecord) && matcher.GetValue(targetRecord) == matcher.GetValue(winningRecord) {
			results = append(results, targetRecord)
			winningRecord = targetRecord
		}
	}

	return results
}

func (recordQuery *recordQuery) RecordsInDateRange(startDate, endDate time.Time) RecordQuery {
	results := []record.Record{}
	for _, record := range recordQuery.records {
		transactionDate := record.TransactionDate
		if (transactionDate.Equal(startDate) || transactionDate.After(startDate)) &&
			(transactionDate.Equal(endDate) || transactionDate.Before(endDate)) {
			results = append(results, record)
		}
	}
	return NewRecordQuery(results)
}

func (recordQuery *recordQuery) RecordsWithDescription(searchString string) RecordQuery {
	records := recordQuery.records
	found := []record.Record{}
	for _, record := range records {
		description := record.TransactionDescription
		searchString = strings.ToUpper(searchString)
		if strings.Contains(description, searchString) {
			found = append(found, record)
		}
	}

	return NewRecordQuery(found)
}

type calculateFunction func(record.Record, *float64)

/*
getDataFromRecords will iterate over all records in assets/finances.json and execute calculate passing each Record as it's argument
*/
func (recordQuery *recordQuery) forEachRecord(calculate calculateFunction) float64 {
	records := recordQuery.records
	var total float64 = 0
	for i := 0; i < len(records); i++ {
		calculate(records[i], &total)
	}
	return total
}
