package app

import (
	"log"
	"net/http"

	"github.com/alufhigi/netServer/db"
	"github.com/urfave/negroni"
)

func New(db *db.DB) *server {
	return &server{Db: db}
}

func (s *server) Run() {
	log.Println("Listening...")
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(&s.router)
	http.ListenAndServe(":5555", n)
}
