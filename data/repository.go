package data

import "github.com/gosum/record"

type Repository interface {
	LoadRecords() (record.Records, error)
}
