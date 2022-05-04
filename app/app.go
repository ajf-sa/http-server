package app

import "net/http"

func New() *server {
	return &server{}
}

func (s *server) Run() {
	http.ListenAndServe(":5555", nil)
}
