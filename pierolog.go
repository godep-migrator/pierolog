package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "pong\n")
	})
	http.ListenAndServe(":9753", nil)
}
