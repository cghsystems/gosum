package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gosum/query"
	"github.com/gosum/record"
)

type link struct {
	Href string `json:"href"`
}

type hyperMedia struct {
	Links map[string]link `json:"_links"`
}

type API struct {
	port        int
	RecordQuery query.RecordQuery `inject:""`
}

type APIHandler interface {
	getHref() string
	handle(http.ResponseWriter, *http.Request)
}

type ResponseMessage struct {
	MetaData MetaData       `json:"metadata"`
	Records  record.Records `json:"data"`
}

type MetaData struct {
	HttpStatus   int    `json:"http_status"`
	ErrorMessage string `json:"error_message"`
}

func NewAPI(port int) *API {
	return &API{
		port: port,
	}
}

func (api *API) Start() {
	http.HandleFunc("/api", api.apiHandler())
	http.HandleFunc("/api/accounts/query/data.json", api.dataHandler())
	port := fmt.Sprintf(":%v", api.port)

	go http.ListenAndServe(port, nil)
}

func (api *API) dataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitValue := r.FormValue("limit")
		limit, err := strconv.Atoi(limitValue)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if err != nil {
			responseMessage := &ResponseMessage{
				MetaData: MetaData{
					HttpStatus:   http.StatusBadRequest,
					ErrorMessage: "Error parsing limit form value: " + err.Error(),
				},
				Records: record.Records{},
			}

			responseJSON, _ := json.Marshal(responseMessage)
			http.Error(w, string(responseJSON), http.StatusBadRequest)
			return
		}

		if limit > 0 {
			w.Write([]byte("[]"))
			return
		}

		responseMessage := &ResponseMessage{
			MetaData: MetaData{
				HttpStatus: http.StatusOK,
			},
			Records: api.RecordQuery.Records(),
		}
		recordsJSON, _ := json.Marshal(responseMessage)
		w.Write(recordsJSON)
	}
}

func (api *API) apiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		links := map[string]link{}
		links["self"] = link{
			Href: "/api",
		}

		hyperMedia := &hyperMedia{
			Links: links,
		}
		hyperMediaJson, _ := json.Marshal(hyperMedia)
		w.Write(hyperMediaJson)
	}
}
