// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * main.go :: main server implementation
// *

package main

import (
	"jsondb"
	//"runtime"
)

func init() {
	//runtime.GOMAXPROCS(4)
}

func main() {
	jsondb.ListenAndServe(jsondb.InitDataStore(), ":12345")
}
