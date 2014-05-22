package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

func main() {
	c := newConfig()

	m := martini.Classic()
	m.Get(`/`, func(w http.ResponseWriter) {
		fd, _ := os.Open(c.Database)
		defer fd.Close()
		w.WriteHeader(200)
		io.Copy(w, fd)
	})
	m.Post(`/`, func(w http.ResponseWriter, r *http.Request) {
		fd, _ := os.Create(c.Database)
		defer fd.Close()
		w.WriteHeader(200)
		io.Copy(fd, r.Body)
	})
	http.ListenAndServe(c.Address, m)
}
