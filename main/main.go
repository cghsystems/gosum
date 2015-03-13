package main

import (
	"github.com/gosum/api"
	"github.com/gosum/data"
	"github.com/gosum/query"
)

func main() {
	repo := data.NewFileRepository("query/assets/test_data.json")
	records, _ := repo.LoadRecords()

	recordQuery := query.NewRecordQuery(records)
	api := api.NewAPI(8080, recordQuery)
	api.Start()
	select {}
}
