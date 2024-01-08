package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		res := []byte("I am an HTTP response")
		_, err := w.Write(res)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occurred"))
		}

	})

	http.ListenAndServe(":8080", nil)
}
