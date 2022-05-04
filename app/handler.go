package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alufhigi/netServer/db"
)

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.notFound(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("index"))
}
func (s *server) about(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("about"))
}
func (s *server) users(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Users"))
}
func (s *server) register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := new(db.User)
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		err = s.Db.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte("User Created"))
		return

	}
	s.notFound(w, r, http.StatusNotFound)

}
