package app

import (
	"log"
	"net/http"

	"github.com/alufhigi/http-server/db"
	"github.com/alufhigi/http-server/utils"
	"github.com/urfave/negroni"
)

func New(db *db.DB) *Server {
	return &Server{Db: db}
}

func (s *Server) Run() {
	log.Println("Listening to " + utils.Config("PORT") + " ...")
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(&s.Router)
	http.ListenAndServe(":"+utils.Config("PORT"), n)
}
