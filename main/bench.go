// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * hdlr.go :: connection handler
// *

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Println(err)
	}
	q := `{"cmd":"has","key":"foo","val":"bar"}`
	t := time.Now().Unix()
	fmt.Println("started timestamp")
	for i := 0; i < 1000000; i++ {
		fmt.Fprintf(conn, q)
	}
	fmt.Println("stoped timestamp")
	ts := time.Now().Unix() - t
	conn.Close()
	fmt.Printf("%d seconds to complete %d requests (%d requests per second)\r\n", ts, 1000000, 1000000/ts)
}
