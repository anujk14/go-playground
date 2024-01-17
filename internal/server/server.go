package server

import (
	"github.com/anujk14/go-playground/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	logger := logger.GetLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler(logger))
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodPost)

	logger.Debug("Starting server.......")

	http.ListenAndServe(":9091", r)
}
