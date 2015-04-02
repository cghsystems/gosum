package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/record"
)

func (api API) QueryAPI(w http.ResponseWriter, r *http.Request) {
	limitValue := r.FormValue("limit")
	limit, err := strconv.Atoi(limitValue)
	log.Debug(fmt.Sprintf("Serving /api/accounts/query/data.json?limit=%v", limit))

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

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
		Records: api.recordQuery.Records(),
	}
	recordsJSON, _ := json.Marshal(responseMessage)
	w.Write(recordsJSON)
}
