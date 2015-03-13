package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gosum/record"
)

func LoadRecords(jsonPath string) (record.Records, error) {
	file, e := ioutil.ReadFile(jsonPath)
	if e != nil {
		return nil, errors.New("Cannot load records from " + jsonPath)
	}

	var records record.Records
	err := json.Unmarshal(file, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
