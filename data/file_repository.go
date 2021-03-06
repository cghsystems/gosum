package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/record"
)

type fileRepository struct {
	jsonPath string
}

func NewFileRepository(jsonPath string) Repository {
	log.Info(fmt.Sprintf("Using datasource %v", jsonPath))
	return &fileRepository{
		jsonPath: jsonPath,
	}
}

func (r *fileRepository) LoadRecords() (record.Records, error) {
	file, e := ioutil.ReadFile(r.jsonPath)
	if e != nil {
		return nil, errors.New("Cannot load records from " + r.jsonPath)
	}

	var records record.Records
	err := json.Unmarshal(file, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
