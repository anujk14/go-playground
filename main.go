package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Simple http handler code is below. This uses DefaultServeMux as the multiplexer. It does not support advanced routing
	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/hello", helloHandler)

	// http.ListenAndServe(":8080", nil)

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodPost)
	http.ListenAndServe(":9091", r)
}
