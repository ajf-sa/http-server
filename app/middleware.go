package app

import "net/http"

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.isAdmin(r) {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

func (s *server) isAdmin(r *http.Request) bool {
	user, ok := r.URL.Query()["user"]
	if !ok || len(user[0]) < 1 {
		return false
	}
	return user[0] == "admin"

}
