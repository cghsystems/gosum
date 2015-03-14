package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cghsystems/gosum/record"
)

type Repository interface {
	LoadRecords() (record.Records, error)
}

func NewHTTPRepository(hosts ...string) Repository {
	return &httpRepository{
		hosts: hosts,
	}
}

type httpRepository struct {
	hosts []string
}

func (r *httpRepository) LoadRecords() (record.Records, error) {
	url := fmt.Sprintf("%v/%s", r.hosts[0], "data")
	response, _ := http.Get(url)

	var records record.Records
	body, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(body, &records)

	return records, nil
}
