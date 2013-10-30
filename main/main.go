package main

import (
	"jsondb"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {
	jsondb.ListenAndServe(jsondb.InitDataStore(), ":12345")
}
