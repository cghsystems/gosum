package api

import (
	"fmt"
	"net/http"
	"time"

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
	http.HandleFunc("/api", requestTimeProxy(api.HyperMedia))
	http.HandleFunc("/api/accounts/query/data.json", requestTimeProxy(api.QueryAPI))
	port := fmt.Sprintf(":%v", api.port)

	log.Info(fmt.Sprintf("Starting API Handler on port %v", port))
	go http.ListenAndServe(port, nil)
}

type httpRequest func(w http.ResponseWriter, r *http.Request)

func requestTimeProxy(execute httpRequest) httpRequest {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		execute(w, r)
		endTime := time.Now()
		requestTime := endTime.Sub(startTime)
		log.Info(fmt.Sprintf("Executed HTTP request in %v ms", requestTime))
		//TODO Record the metric
	}
}
