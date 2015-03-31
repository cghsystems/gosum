package main

import (
	"flag"

	"github.com/cghsystems/gosum/api"
	"github.com/cghsystems/gosum/data"
	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/query"
)

var datasource = flag.String(
	"datasource",
	"query/assets/test_data.json",
	"Location of the source finances data file")

func init() {
	log.Init()
	log.Info("Starting gosum...")
}

func main() {
	flag.Parse()

	repo := data.NewFileRepository(*datasource)
	records, _ := repo.LoadRecords()

	recordQuery := query.NewRecordQuery(records)
	api := api.NewAPI(8080, recordQuery)
	api.Start()
	select {}
}
