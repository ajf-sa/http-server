package app

func (s *Server) Routers() {
	s.Router.HandleFunc("/about", Chain(s.About, s.Method("GET")))
	s.Router.HandleFunc("/register", Chain(s.Register, s.Method("POST")))
	s.Router.HandleFunc("/login", Chain(s.Login, s.Method("POST")))
	s.Router.HandleFunc("/users", Chain(s.Users, s.Method("GET"), s.AdminOnly(), s.LoginOnly()))
	s.Router.HandleFunc("/", Chain(s.Index, s.Method("GET")))
}
