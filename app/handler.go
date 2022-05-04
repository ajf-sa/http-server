package app

import "net/http"

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func (s *server) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
func (s *server) users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))
}
