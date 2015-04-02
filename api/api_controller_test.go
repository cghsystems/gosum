package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cghsystems/gosum/api"
	"github.com/cghsystems/gosum/query"
	"github.com/cghsystems/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTTP API", func() {
	const port = 9898

	var (
		httpAPI      *api.API
		actualRecord record.Record
	)

	BeforeSuite(func() {
		actualRecord = record.Record{SortCode: "test"}
		query := query.NewRecordQuery(record.Records{actualRecord})

		httpAPI = api.NewAPI(port, query)
		httpAPI.Start()
	})

	Context(".Start", func() {
		Context("/api", func() {
			It("returns 200 OK status code", func() {
				resp, _ := http.Get("http://localhost:9898/api")
				Ω(resp.StatusCode).Should(Equal(200))
			})

			It("returns the hypermedia", func() {
				expectedJson := ` {"_links":{"self":{"href":"/api"}}} `
				resp, _ := http.Get("http://localhost:9898/api")
				body, _ := ioutil.ReadAll(resp.Body)
				Ω(string(body)).Should(MatchJSON(expectedJson))
			})
		})

		Context("/api/accounts/query/data.json?limit=0", func() {
			var responseMessage api.ResponseMessage

			BeforeEach(func() {
				resp, err := http.Get("http://localhost:9898/api/accounts/query/data.json?limit=0")
				Ω(err).ShouldNot(HaveOccurred())
				body, err := ioutil.ReadAll(resp.Body)
				Ω(err).ShouldNot(HaveOccurred())
				json.Unmarshal(body, &responseMessage)
			})

			It("contains the expected meta data", func() {
				Ω(http.StatusOK).Should(Equal(responseMessage.MetaData.HttpStatus))
			})
		})
	})
})
