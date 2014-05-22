package server

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/meatballhat/pierolog/notifications"
)

type server struct {
	cfg      *config
	notifier *notifications.HipChatNotifier
}

func Main() {
	newServer().Run()
}

func newServer() *server {
	cfg := newConfig()
	return &server{
		cfg: cfg,
		notifier: notifications.NewHipChatNotifier(
			cfg.HipChatAuthToken,
			cfg.HipChatRoomID,
			cfg.HipChatFrom,
		),
	}
}

func (srv *server) Run() {
	m := martini.Classic()

	m.Get(`/`, srv.handleGetAll)
	m.Post(`/`, srv.handleCreate)

	os.Setenv("PORT", srv.cfg.Port)

	log.Printf("running with config: %#v\n", srv.cfg)
	m.Run()
}

func (srv *server) handleGetAll(w http.ResponseWriter) {
	fd, _ := os.Open(srv.cfg.Database)
	defer fd.Close()
	w.WriteHeader(200)
	io.Copy(w, fd)
}

func (srv *server) handleCreate(w http.ResponseWriter, r *http.Request) {
	fd, _ := os.Create(srv.cfg.Database)
	defer fd.Close()
	w.WriteHeader(200)
	io.Copy(fd, r.Body)
	err := srv.notifier.Notify("ermahgerd database is updated")
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
