package app

import (
	"net/http"

	"github.com/alufhigi/http-server/db"
)

type Server struct {
	Db     *db.DB
	Router http.ServeMux
}
