package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HelloResponse struct {
	Name     string
	Greeting string
}

func rootHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("Received root request")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		res := []byte("I am an HTTP response")
		_, err := w.Write(res)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occurred"))
		}
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	name := req.Header.Get("name")
	if name == "" {
		name = "Random dude"
	}

	greeting := "Hello!"

	h := HelloResponse{
		Name:     name,
		Greeting: greeting,
	}

	w.Header().Set("content-type", "application/json")

	err := json.NewEncoder(w).Encode(h)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred"))
	}
}
