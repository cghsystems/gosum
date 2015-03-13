package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gosum/record"
)

type fileRepository struct {
	jsonPath string
}

func NewFileRepository(jsonPath string) Repository {
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
