// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * hdlr.go :: connection handler
// *

package jsondb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn, ds *DataStore) {
	log.Printf("%v connected\n", conn.RemoteAddr())
	dec := json.NewDecoder(conn)
	for {
		var r Request
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			conn.Close()
			return
		}
		switch r.Cmd {
		case "has":
			fmt.Fprintf(conn, "%s\r\n", ds.has(&r).EncodeResp())
			break
		case "add":
			fmt.Fprintf(conn, "%s\r\n", ds.add(&r).EncodeResp())
			break
		case "set":
			fmt.Fprintf(conn, "%s\r\n", ds.set(&r).EncodeResp())
			break
		case "get":
			fmt.Fprintf(conn, "%s\r\n", ds.get(&r).EncodeResp())
			break
		case "del":
			fmt.Fprintf(conn, "%s\r\n", ds.del(&r).EncodeResp())
			break
		default:
			fmt.Fprintf(conn, "%s\r\n", ds.err(&r).EncodeResp())
			break
		}
	}
}
