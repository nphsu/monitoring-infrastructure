package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	return &Server{
		router: newRouter(),
	}, nil
}

func (s *Server) Run(port int) {
	p := strconv.Itoa(port)
	s.router.Run(":" + p)
}
