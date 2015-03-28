package db_test

import (
	"time"

	"github.com/cghsystems/gosum/db"
	"github.com/cghsystems/gosum/record"
	"github.com/fzzy/radix/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	const redisURL = "127.0.0.1:6379"

	var testRecord = record.Record{
		TransactionType:        record.CREDIT,
		SortCode:               "12-34-56",
		AccountNumber:          "123456789",
		TransactionDescription: "A Test Record",
		DebitAmount:            12.12,
		CreditAmount:           0.0,
		Balance:                12.12,
	}

	var (
		client *db.Client
		err    error
	)

	BeforeSuite(func() {
		startRedis()
	})

	cleanRedis := func() {
		c, _ := redis.DialTimeout("tcp", redisURL, time.Duration(10)*time.Second)
		defer c.Close()
		c.Cmd("select", 0)
		c.Cmd("FLUSHDB")
	}

	BeforeEach(func() {
		client = db.New(redisURL)
		cleanRedis()
	})

	AfterEach(func() {
		client.Close()
	})

	Context("close", func() {
		It("closes the connection", func() {
			client.Close()
			err := client.Set(testRecord)
			Ω(err).Should(HaveOccurred())
		})
	})

	Context(".Set", func() {
		BeforeEach(func() {
			err = client.Set(testRecord)
		})

		It("does not return error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("sets the expected data", func() {
			actualRecord := client.Get(testRecord.ID())
			Ω(testRecord.Balance).Should(Equal(actualRecord.Balance))
		})
	})

	Context("BulkSet", func() {
		BeforeEach(func() {
			records := record.Records{testRecord}
			err = client.BulkSet(records)
		})

		It("does not return error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
