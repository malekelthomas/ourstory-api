package server

import "github.com/gin-gonic/gin"

type server struct {
	Router *gin.Engine
}

func NewServer() *server {
	r := gin.Default()

	return &server{Router: r}
}

type ServerErrorMsg struct {
	Message string
	Error   string
}

func (s server) Run(port string) {
	s.Router.Run(port)
}
