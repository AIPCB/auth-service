package server

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/register", s.RegisterHandler()).Methods("POST")

}
