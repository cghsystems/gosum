package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/facebookgo/inject"
	"github.com/gosum/api"
	"github.com/gosum/query"
	"github.com/gosum/record"
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

		httpAPI = api.NewAPI(port)

		if err := inject.Populate(query, httpAPI); err != nil {
			Ω(err).ShouldNot(HaveOccurred())
		}

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

			It("contains no error message ", func() {
				Ω(responseMessage.MetaData.ErrorMessage).Should(BeEmpty())
			})

			It("contains the expected number of records", func() {
				Ω(len(responseMessage.Records)).Should(Equal(1))
			})

			It("contains the expected record", func() {
				record, _ := responseMessage.Records.First()
				Ω(record.SortCode).Should(Equal(actualRecord.SortCode))
			})

			It("returns no data if limit is > 0", func() {
				expectedJson := `[]`
				resp, err := http.Get("http://localhost:9898/api/accounts/query/data.json?limit=1")
				Ω(err).NotTo(HaveOccurred())
				body, _ := ioutil.ReadAll(resp.Body)
				Ω(string(body)).Should(MatchJSON(expectedJson))
			})
		})

		Context("/api/accounts/query/data.json?limit=unparseable", func() {
			var (
				resp            *http.Response
				responseMessage api.ResponseMessage
			)

			BeforeEach(func() {
				var err error
				resp, err = http.Get("http://localhost:9898/api/accounts/query/data.json?limit=wibble1")
				Ω(err).NotTo(HaveOccurred())
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(bodyBytes, &responseMessage)
			})

			It("contains zero records", func() {
				Ω(len(responseMessage.Records)).Should(Equal(0))
			})

			It("metadata contains bad request status code", func() {
				httpStatusCode := responseMessage.MetaData.HttpStatus
				Ω(http.StatusBadRequest).Should(Equal(httpStatusCode))
			})

			It("metadata contains the error message", func() {
				errorMessage := responseMessage.MetaData.ErrorMessage
				expectedError := `Error parsing limit form value: strconv.ParseInt: parsing "wibble1": invalid syntax`
				Ω(errorMessage).Should(Equal(expectedError))
			})

			It("returns a 400 in the header", func() {
				Ω(http.StatusBadRequest).Should(Equal(resp.StatusCode))
			})
		})
	})
})
