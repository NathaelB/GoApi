package main

func (s *Server) Routes() {
	s.engine.GET("/peoples", s.getPeoples)
}
