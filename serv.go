// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * serv.go :: server
// *

package jsondb

import (
	"log"
	"net"
)

func ListenAndServe(ds *DataStore, port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %v\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("*accept: %s\n", err)
			continue
		}
		go handleConn(conn, ds)
	}
}
