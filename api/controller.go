package api

import (
	"fmt"
	"net/http"

	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/query"
)

func NewAPI(port int, recordQuery query.RecordQuery) *API {
	return &API{
		port:        port,
		recordQuery: recordQuery,
	}
}

func (api *API) Start() {
	http.HandleFunc("/api", api.HyperMedia)
	http.HandleFunc("/api/accounts/query/data.json", api.QueryAPI)
	port := fmt.Sprintf(":%v", api.port)

	log.Info(fmt.Sprintf("Starting API Handler on port %v", port))
	go http.ListenAndServe(port, nil)
}
