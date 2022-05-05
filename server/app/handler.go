package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/alufhigi/http-server/db"
	"github.com/alufhigi/http-server/utils"
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
	p := new(db.Pagination)
	limit, ok := r.URL.Query()["limit"]
	if ok {
		p.Limit, _ = strconv.Atoi(limit[0])
	} else {
		p.Limit = 10
	}

	page, ok := r.URL.Query()["page"]
	if ok {
		pp, _ := strconv.Atoi(page[0])
		if pp <= 1 {
			p.Page = 0
		} else {
			p.Page = pp - 1
		}
	} else {
		p.Page = 0
	}
	u, err := s.Db.FindAllUser(p)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(u)
	w.Write([]byte(b))
}
func (s *server) register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := new(db.User)
	err = user.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pss, er := user.Password.Hash()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = pss
	err = s.Db.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User Created"))
	return

}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := new(db.User)
	err = user.UnmarshalJSON(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(user)
	if err = user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u, err := s.Db.FindOneUserByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if u.Password.Check(string(user.Password)) {
		log.Println("Login Success")
		token, _ := utils.CreateToken(u.ID)
		b, _ := json.Marshal(token)
		w.Write([]byte(b))
		return

	}
	w.Write([]byte("Login Failed"))
}
