package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine
	store  Store
}

func newServer() *Server {
	s := &Server{
		engine: gin.Default(),
	}

	s.Routes()
	return s
}

func main() {
	s := newServer()
	s.store = &dbStore{}
	err := s.store.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer s.store.Disconnect()

	s.engine.Run(":3333")
}
