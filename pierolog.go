package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get(`/`, func(w http.ResponseWriter) {
		fd, _ := os.Open("pierolog.db")
		defer fd.Close()
		w.WriteHeader(200)
		io.Copy(w, fd)
	})
	m.Post(`/`, func(w http.ResponseWriter, r *http.Request) {
		fd, _ := os.Create("pierolog.db")
		defer fd.Close()
		w.WriteHeader(200)
		io.Copy(fd, r.Body)
	})
	http.ListenAndServe(":9753", m)
}
