// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * hdlr.go :: connection handler
// *

package jsondb

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn, ds *DataStore) {
	log.Printf("%v connected\n", conn.RemoteAddr())
	dec, enc := json.NewDecoder(conn), json.NewEncoder(conn)
	for {
		var req Request
		if err := dec.Decode(&req); err == io.EOF {
			break
		} else if err != nil {
			conn.Close()
			return
		} else {
			conn.SetDeadline(time.Now().Add(time.Duration(5) * time.Minute))
		}
		switch req.Cmd {
		case "exp":
			enc.Encode(ds.exp(&req))
		case "ttl":
			enc.Encode(ds.ttl(&req))
		case "has":
			enc.Encode(ds.has(&req))
		case "add":
			enc.Encode(ds.add(&req))
		case "set":
			enc.Encode(ds.set(&req))
		case "get":
			enc.Encode(ds.get(&req))
		case "del":
			enc.Encode(ds.del(&req))
		default:
			enc.Encode(ds.err(&req))
		}
	}
}
