package main

import (
	"fmt"
	"net/http"
)

func handler(write http.ResponseWriter, request *http.Request) {
	fmt.Printf("Hello %v", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:1337", nil)
}
