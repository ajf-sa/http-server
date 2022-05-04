package app

import (
	"net/http"

	"github.com/alufhigi/netServer/db"
)

type server struct {
	Db     *db.DB
	router http.ServeMux
}
