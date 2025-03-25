package server

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/register", s.RegisterHandler()).Methods("POST")
	s.router.HandleFunc("/login", s.LoginHandler()).Methods("POST")
}
