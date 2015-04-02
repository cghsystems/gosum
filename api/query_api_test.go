package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/cghsystems/gosum/api"
	"github.com/cghsystems/gosum/query"
	"github.com/cghsystems/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("QueryApi", func() {
	var (
		writer          *httptest.ResponseRecorder
		request         *http.Request
		responseMessage api.ResponseMessage
		actualRecord    record.Record
		err             error
	)

	JustBeforeEach(func() {
		Ω(err).ToNot(HaveOccurred())

		actualRecord = record.Record{SortCode: "test"}
		query := query.NewRecordQuery(record.Records{actualRecord})
		writer = httptest.NewRecorder()
		api := api.NewAPI(80, query)
		api.QueryAPI(writer, request)

		body, err := ioutil.ReadAll(writer.Body)
		Ω(err).ShouldNot(HaveOccurred())
		json.Unmarshal(body, &responseMessage)
	})

	Context("query with zero limit", func() {
		BeforeEach(func() {
			request, err = http.NewRequest("GET", "query/data.json?limit=0", nil)
		})

		It("returns http OK status", func() {
			Ω(writer.Code).To(Equal(http.StatusOK))
		})

		It("has a content type of application/json", func() {
			contentType := writer.HeaderMap["Content-Type"][0]
			Ω(contentType).To(Equal("application/json"))
		})

		It("has the expected number of records in the JSON payload", func() {
			Ω(len(responseMessage.Records)).Should(Equal(1))
		})

		It("contains the expected record", func() {
			record, _ := responseMessage.Records.First()
			Ω(record.SortCode).Should(Equal(actualRecord.SortCode))
		})

		It("contains the expected meta data", func() {
			Ω(http.StatusOK).Should(Equal(responseMessage.MetaData.HttpStatus))
		})

		It("contains no error message ", func() {
			Ω(responseMessage.MetaData.ErrorMessage).Should(BeEmpty())
		})
	})

	Context("query with limit of greater that 0", func() {
		BeforeEach(func() {
			request, err = http.NewRequest("GET", "query/data.json?limit=1", nil)
		})

		It("returns no data if limit is > 0", func() {
			Ω(len(responseMessage.Records)).Should(Equal(1))
		})
	})

	Context("limit is unparseable", func() {
		BeforeEach(func() {
			request, err = http.NewRequest("GET", "query/data.json?limit=wibble1", nil)
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
			Ω(http.StatusBadRequest).Should(Equal(writer.Code))
		})
	})
})
