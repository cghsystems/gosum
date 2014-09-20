package record

import (
	"errors"
	"time"
)

type Records []Record

func (records Records) First() (Record, error) {
	if len(records) < 1 {
		return NilRecord(), errors.New("There are no records to find")
	}
	return records[0], nil
}

func (record Records) DeleteAll() Records {
	return Records{}
}

func (records Records) Delete(record Record) (Records, error) {
	if len(records) < 1 {
		return nil, errors.New("There are no records to Delete")
	}

	length := len(records)
	toDeleteIndex := records.indexOf(record)

	if toDeleteIndex == -1 {
		return nil, errors.New("Cannot find record to delete")
	}

	records[toDeleteIndex] = records[length-1]
	records = records[:length-1]
	return records, nil
}

func (records Records) indexOf(target Record) int {
	for index, record := range records {
		if record == target {
			return index
		}
	}
	return -1
}

// Record is the struct that represents the a record taken from the collection in assests/finances.json
type Record struct {
	TransactionDate        time.Time `json:"transaction_date"`
	TransactionType        string    `json:"transaction_type"`
	SortCode               string    `json:"sort_code"`
	AccountNumber          string    `json:"account_number"`
	TransactionDescription string    `json:"transaction_description"`
	DebitAmount            float64   `json:"debit_amount"`
	CreditAmount           float64   `json:"credit_amount"`
	Balance                float64   `json:"balance"`
}

func (record Record) IsNilRecord() bool {
	return record == NilRecord() || record == Record{}
}

func NilRecord() Record {
	return Record{
		TransactionDate:        time.Time{},
		TransactionType:        "nil",
		SortCode:               "nil",
		AccountNumber:          "nil",
		TransactionDescription: "nil",
		DebitAmount:            float64(-99999999),
		CreditAmount:           float64(-99999999),
		Balance:                float64(-99999999),
	}
}
