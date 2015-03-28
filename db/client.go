package db

import (
	"encoding/json"
	"time"

	"github.com/cghsystems/gosum/record"
	"github.com/fzzy/radix/redis"
)

type Client struct {
	redis *redis.Client
}

func New(url string) *Client {
	c, _ := redis.DialTimeout("tcp", url, time.Duration(10)*time.Second)
	c.Cmd("select", 0)
	return &Client{redis: c}
}

func (c *Client) BulkSet(records record.Records) error {
	for _, record := range records {
		c.Set(record)
	}
	return nil
}

func (c *Client) Set(record record.Record) error {
	json, err := json.Marshal(record)

	if err != nil {
		return err
	}

	r := c.redis.Cmd("set", record.ID(), json)
	return r.Err
}

func (c *Client) Close() {
	c.redis.Close()
}

func (c *Client) Get(key string) record.Record {
	bytes, _ := c.redis.Cmd("get", key).Bytes()
	var record record.Record
	_ = json.Unmarshal(bytes, &record)
	return record
}
