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
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to %v", url)
	}

	var records record.Records
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unexpected response from server", err.Error())
	}

	err = json.Unmarshal(body, &records)
	if err != nil {
		return nil, fmt.Errorf("unexpected response from server: %v", err.Error())
	}

	return records, nil
}
