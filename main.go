package main

import "net/http"

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
	res := []byte("Hello there!")
	_, err := w.Write(res)

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
