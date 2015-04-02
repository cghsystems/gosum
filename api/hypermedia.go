package api

import (
	"encoding/json"
	"net/http"

	"github.com/cghsystems/gosum/query"
)

type Link struct {
	Href string `json:"href"`
}

type HyperMedia struct {
	Links map[string]Link `json:"_links"`
}

type API struct {
	port        int
	recordQuery query.RecordQuery
}

func (api *API) HyperMedia(w http.ResponseWriter, r *http.Request) {
	links := map[string]Link{}
	links["self"] = Link{
		Href: "/api",
	}

	hyperMedia := &HyperMedia{
		Links: links,
	}
	hyperMediaJson, _ := json.Marshal(hyperMedia)
	w.Write(hyperMediaJson)
}
