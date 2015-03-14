package data

import "github.com/cghsystems/gosum/record"

type Repository interface {
	LoadRecords() (record.Records, error)
}
