package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Payload struct {
	Payload Data
}

type Data struct {
	Book, Podcast map[string]string
}

func handler(write http.ResponseWriter, request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	response, _ := getJSONResponse()

	fmt.Fprintf(write, string(response))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:1337", nil)
}

func getJSONResponse() ([]byte, error) {
	books := make(map[string]string)

	books["Ego is the enemy"] = "Get rid of your ego and achieve greater things"
	books["Manage Your Day-to-Day"] = "Great book for building daily routines"

	podcasts := make(map[string]string)

	podcasts["seanwes podcast"] = "How to pursuit and make a living of your passion"
	podcasts["Changelog"] = "Interviews with the most influential people from the open source community"

	data := Data{books, podcasts}
	payload := Payload{data}

	return json.MarshalIndent(payload, "", " ")
}
