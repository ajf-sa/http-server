package app

import (
	"log"
	"net/http"

	"github.com/alufhigi/http-server/db"
	"github.com/urfave/negroni"
)

func New(c Config) (*Server, error) {
	db, err := db.New(c.DB)
	if err != nil {
		return nil, err
	}

	return &Server{Db: db, Config: c}, nil
}

func (s *Server) Run() {
	log.Println("Listening to " + s.Config.Port + " ...")
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(&s.Router)
	http.ListenAndServe(":"+s.Config.Port, n)
}

func (s *Server) CloseDB() {
	s.Db.Db.Close()
}
