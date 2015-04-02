package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/record"
)

type ResponseMessage struct {
	MetaData MetaData       `json:"metadata"`
	Records  record.Records `json:"data"`
}

type MetaData struct {
	HttpStatus   int    `json:"http_status"`
	ErrorMessage string `json:"error_message"`
}

func (api API) QueryAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	limitValue := r.FormValue("limit")
	limit, err := strconv.Atoi(limitValue)

	log.Debug(fmt.Sprintf("Serving /api/accounts/query/data.json?limit=%v", limit))

	if err != nil {
		handleLimitValueConversionError(err, w)
		return
	}

	if limit > 0 {
		w.Write([]byte("[]"))
		return
	}
	api.writeSuccessResponse(w)
}

func (api API) writeSuccessResponse(w http.ResponseWriter) {
	responseMessage := &ResponseMessage{
		MetaData: MetaData{
			HttpStatus: http.StatusOK,
		},
		Records: api.recordQuery.Records(),
	}
	recordsJSON, _ := json.Marshal(responseMessage)
	w.Write(recordsJSON)
}

func handleLimitValueConversionError(err error, w http.ResponseWriter) {
	responseMessage := &ResponseMessage{
		MetaData: MetaData{
			HttpStatus:   http.StatusBadRequest,
			ErrorMessage: "Error parsing limit form value: " + err.Error(),
		},
		Records: record.Records{},
	}

	responseJSON, _ := json.Marshal(responseMessage)
	http.Error(w, string(responseJSON), http.StatusBadRequest)
}
