package app

func (s *server) Routers() {
	s.router.HandleFunc("/about", Chain(s.about, s.method("GET")))
	s.router.HandleFunc("/register", Chain(s.register, s.method("POST")))
	s.router.HandleFunc("/login", Chain(s.login, s.method("POST")))
	s.router.HandleFunc("/users", Chain(s.users, s.method("GET"), s.adminOnly(), s.loginOnly()))
	s.router.HandleFunc("/", Chain(s.index, s.method("GET")))
}
