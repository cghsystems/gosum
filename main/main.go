package main

import (
	"github.com/cghsystems/gosum/api"
	"github.com/cghsystems/gosum/data"
	"github.com/cghsystems/gosum/query"
)

func main() {
	repo := data.NewFileRepository("query/assets/test_data.json")
	records, _ := repo.LoadRecords()

	recordQuery := query.NewRecordQuery(records)
	api := api.NewAPI(8080, recordQuery)
	api.Start()
	select {}
}
