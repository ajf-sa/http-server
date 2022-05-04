package app

import (
	"net/http"
)

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.isAdmin(r) {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

func (s *server) loginOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.isLogin(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
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
	//TODO check user is admin
	return user[0] == "admin"

}

func (s *server) isLogin(r *http.Request) bool {
	token, ok := r.Header["Authorization"]
	if !ok || len(token[0]) < 1 {
		return false
	}
	token[0] = token[0][7:]
	//TODO check token is valid
	return token[0] == "1234"

}

func (s *server) notFound(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		w.Write([]byte("404 - Page not found"))
	}

}
