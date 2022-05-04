package app

func (s *server) Routers() {

	s.router.HandleFunc("/about", s.about)
	s.router.HandleFunc("/register", s.register)
	s.router.HandleFunc("/users", s.loginOnly(s.adminOnly(s.users)))
	s.router.HandleFunc("/", s.index)
}
