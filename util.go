// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * util.go :: utilities
// *

package jsondb

import (
	"encoding/json"
	"log"
)

type Request struct {
	Cmd string      `json:"cmd"`
	Key string      `json:"key"`
	Val interface{} `json:"val"`
}

type Response map[string]interface{}

func (r Response) EncodeResp() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	return b
}
