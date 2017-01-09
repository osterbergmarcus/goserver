package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func prettyfyRequest(req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func handler(write http.ResponseWriter, request *http.Request) {
	prettyfyRequest(request)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:1337", nil)
}
