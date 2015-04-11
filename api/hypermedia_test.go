package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/cghsystems/gosum/api"
	"github.com/cghsystems/gosum/log"
	"github.com/cghsystems/gosum/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hypermedia", func() {
	var (
		response        *httptest.ResponseRecorder
		request         *http.Request
		responseMessage *api.HyperMedia
		err             error
	)
	JustBeforeEach(func() {
		log.Init()
		err := metrics.InitFileRecorder("/tmp/test_metrics")

		Ω(err).ToNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())

		response = httptest.NewRecorder()
		api := api.NewAPI(80, nil)
		api.HyperMedia(response, request)

		body, err := ioutil.ReadAll(response.Body)
		Ω(err).ShouldNot(HaveOccurred())
		json.Unmarshal(body, &responseMessage)
	})

	BeforeEach(func() {
		request, err = http.NewRequest("GET", "query/data.json?limit=0", nil)
	})

	It("returns 200 OK status code", func() {
		resp, _ := http.Get("http://localhost:9898/api")
		Ω(resp.StatusCode).Should(Equal(200))
	})

	It("contains the api link", func() {
		apiLink := api.Link{"/api"}
		Ω(responseMessage.Links).Should(ContainElement(apiLink))
	})
})
