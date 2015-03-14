package data_test

import (
	"net/http"

	"github.com/cghsystems/gosum/data"
	"github.com/cghsystems/gosum/record"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("HttpRepository", func() {
	var (
		repository data.Repository
		server     *ghttp.Server
		records    record.Records
		err        error
	)

	Context("running server", func() {
		BeforeEach(func() {
			server = ghttp.NewServer()
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/data"),
					ghttp.RespondWith(http.StatusOK, `[
			      {"transaction_type": "CPT"}	
				  ]`),
				),
			)

			repository = data.NewHTTPRepository(server.URL())
			records, err = repository.LoadRecords()
		})

		AfterEach(func() {
			server.Close()
		})

		It("GETS from the /data endpoint", func() {
			Ω(server.ReceivedRequests()).Should(HaveLen(1))
		})

		It("returns the expected records", func() {
			Ω(records).To(HaveLen(1))
		})
	})
})
