package main

import (
	"github.com/gosum/api"
	"github.com/gosum/data"
	"github.com/gosum/query"
)

func main() {
	records, _ := data.LoadRecords("query/assets/test_data.json")
	recordQuery := query.NewRecordQuery(records)
	api := api.NewAPI(8080)
	api.RecordQuery = recordQuery
	api.Start()
	select {}
}
