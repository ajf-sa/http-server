package app

import (
	"log"
	"net/http"

	"github.com/alufhigi/netServer/db"
	"github.com/alufhigi/netServer/utils"
	"github.com/urfave/negroni"
)

func New(db *db.DB) *server {
	return &server{Db: db}
}

func (s *server) Run() {
	log.Println("Listening to " + utils.Config("PORT") + " ...")
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(&s.router)
	http.ListenAndServe(":"+utils.Config("PORT"), n)
}
