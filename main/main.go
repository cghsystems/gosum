package main

import (
	"github.com/gosum/api"
	"github.com/gosum/query"
	"github.com/gosum/repository"
)

func main() {
	records, _ := repository.LoadRecords("query/assets/test_data.json")
	recordQuery := query.NewRecordQuery(records)
	api := api.NewAPI(8080)
	api.RecordQuery = recordQuery
	api.Start()
	select {}
}
