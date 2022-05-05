package app

import (
	"net/http"

	"github.com/alufhigi/http-server/db"
)

type server struct {
	Db     *db.DB
	router http.ServeMux
}
