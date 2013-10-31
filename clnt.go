// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * clnt.go :: client
// *

package jsondb

import (
	"encoding/json"
	"log"
	"net"
)

type Client struct {
	host string
	resp interface{}
	conn net.Conn
	enc  *json.Encoder
	dec  *json.Decoder
	tx   bool
}

func InitClient(host string) *Client {
	return &Client{
		host: host,
	}
}

func (self *Client) NewTx() {
	conn, err := net.Dial("tcp", self.host)
	if err != nil {
		log.Println(err)
	}
	self.conn = conn
	self.enc = json.NewEncoder(self.conn)
	self.dec = json.NewDecoder(self.conn)
	self.tx = true
}

func (self *Client) EndTx() {
	err := self.conn.Close()
	if err != nil {
		log.Println(err)
	}
	self.tx = false
}

func (self *Client) InTx() bool {
	return self.tx
}

func (self *Client) Raw(json string) {
	self.enc.Encode(json)
}
