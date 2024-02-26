package main

import (
	"log"
	"net/http"
)

var (
	counter    = 0
	firstHost  = "http://localhost:8081"
	secondHost = "http://localhost:8082"
)

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	address := r.URL.String()
	if counter == 0 {
		counter++
		http.NewRequest(r.Method, firstHost+address, r.Body)
	}

	if counter == 1 {
		counter--
		http.NewRequest(r.Method, secondHost+address, r.Body)
	}

}
