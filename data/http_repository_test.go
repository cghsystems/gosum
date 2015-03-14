package data_test

import (
	"fmt"
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
		})

		AfterEach(func() {
			server.Close()
		})

		Context("functioning server", func() {
			BeforeEach(func() {
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

			It("GETS from the /data endpoint", func() {
				Ω(server.ReceivedRequests()).Should(HaveLen(1))
			})

			It("returns the expected records", func() {
				Ω(records).To(HaveLen(1))
			})

			It("does not return error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("server returns unexpected json", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/data"),
						ghttp.RespondWith(http.StatusOK, `{"something": "else"}`),
					),
				)
				repository = data.NewHTTPRepository(server.URL())
				records, err = repository.LoadRecords()
			})

			It("returns an error", func() {
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("server returns no body", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/data"),
					),
				)
				repository = data.NewHTTPRepository(server.URL())
				records, err = repository.LoadRecords()
			})

			It("returns an error", func() {
				Ω(err).To(MatchError("unexpected response from server: unexpected end of JSON input"))
			})
		})
	})

	Context("stopped server", func() {
		var url string
		BeforeEach(func() {
			url = "http://127.0.0.1/some/url"
			repository = data.NewHTTPRepository(url)
			records, err = repository.LoadRecords()
		})

		It("returns error", func() {
			Ω(err).To(MatchError(fmt.Sprintf("cannot connect to %v/data", url)))
		})
	})
})
