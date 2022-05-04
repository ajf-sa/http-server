package app

import "net/http"

func (s *server) Routers() {
	http.HandleFunc("/", s.index)
	http.HandleFunc("/about", s.about)
	http.HandleFunc("/users", s.adminOnly(s.users))

}
