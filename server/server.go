package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/meatballhat/pierolog/notifications"
)

type server struct {
	cfg      *config
	notifier *notifications.HipChatNotifier
	db       *redis.Pool
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Main is the whole shebang!
func Main() {
	server, err := newServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}

func newServer() (*server, error) {
	cfg := newConfig()
	db := newPool(cfg.RedisURL, cfg.RedisPassword)

	return &server{
		cfg: cfg,
		notifier: notifications.NewHipChatNotifier(
			cfg.HipChatAuthToken,
			cfg.HipChatRoomID,
			cfg.HipChatFrom,
		),
		db: db,
	}, nil
}

func (srv *server) Run() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get(`/`, srv.handleGetAll)
	m.Post(`/`, srv.handleCreate)

	os.Setenv("PORT", srv.cfg.Port)

	log.Printf("running with config: %#v\n", srv.cfg)
	m.Run()
}

func (srv *server) handleGetAll(w http.ResponseWriter, r render.Render) {
	conn := srv.db.Get()
	defer conn.Close()

	entries, err := redis.Strings(conn.Do("SMEMBERS", "pierolog:entries"))
	if err != nil {
		r.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	log.Printf("Raw entries: %#v\n", entries)

	r.JSON(200, map[string]interface{}{"entries": entries})
}

func (srv *server) handleCreate(w http.ResponseWriter, r *http.Request, rnd render.Render) {
	entryBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rnd.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	conn := srv.db.Get()
	defer conn.Close()

	entryString := string(entryBytes)
	log.Printf("Adding entry %#v\n", entryString)
	_, err = conn.Do("SADD", "pierolog:entries", entryString)
	if err != nil {
		rnd.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	rnd.JSON(201, map[string]string{"ok": "yup"})

	err = srv.notifier.Notify("ermahgerd database is updated")
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
