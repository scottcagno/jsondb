// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * bench.go :: benchmarker
// *

package main

import (
	"encoding/json"
	"fmt"
	"jsondb"
	"log"
	"net"
	//"runtime"
	"time"
)

func init() {
	//runtime.GOMAXPROCS(4)
}

const N = 1000000 // 1 million

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Println(err)
	}
	dec, enc := json.NewDecoder(conn), json.NewEncoder(conn)
	t := time.Now().Unix()
	for i := 0; i < N; i++ {
		enc.Encode(jsondb.Request{"set", fmt.Sprintf("%d", i), "bar"})
		var res int64
		dec.Decode(&res)
		if res < 0 {
			log.Panic("empty response!\n")
		}
	}
	ts := time.Now().Unix() - t
	conn.Close()
	fmt.Printf("%d seconds to complete %d requests (%d requests per second)\r\n", ts, N, N/ts)
}
