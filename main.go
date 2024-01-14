package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logger := getLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler(logger))
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodPost)

	logger.Debug("Starting server.......")

	http.ListenAndServe(":9091", r)
}
