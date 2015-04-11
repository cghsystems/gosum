package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/metrics"
	"github.com/cghsystems/gosum/query"
)

func NewAPI(port int, recordQuery query.RecordQuery) *API {
	return &API{
		port:        port,
		recordQuery: recordQuery,
	}
}

func (api *API) Start() {
	registerHandleFunc("/api", api.HyperMedia)
	registerHandleFunc("/api/accounts/query/data.json", api.QueryAPI)
	port := fmt.Sprintf(":%v", api.port)

	log.Info(fmt.Sprintf("Starting API Handler on port %v", port))
	go http.ListenAndServe(port, nil)
}

type httpRequest func(w http.ResponseWriter, r *http.Request)

func registerHandleFunc(pattern string, request httpRequest) {
	http.HandleFunc(pattern, requestTimeProxy(pattern, request))
}

func requestTimeProxy(url string, execute httpRequest) httpRequest {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		execute(w, r)
		endTime := time.Now()
		metrics.Record(url, startTime, endTime)
	}
}
