package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getPeoples(c *gin.Context) {
	people, err := s.store.GetPeoples()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}
