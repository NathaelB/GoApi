package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getPeoples(c *gin.Context) {
	people, err := s.GetPeoples()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}
