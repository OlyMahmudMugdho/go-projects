package api

import (
	"net/http"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/services/hello"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	helloHandler := hello.NewHello()
	helloHandler.RegisterRoutes(router)

	return http.ListenAndServe(":"+s.port, router)
}
