package main

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Name     string
	Greeting string
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	res := []byte("I am an HTTP response")
	_, err := w.Write(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred"))
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	name := req.Header.Get("name")
	if name == "" {
		name = "Random dude"
	}

	greeting := "Hello!"

	// res := []byte("Hello there!")
	// _, err := w.Write(res)

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

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}
