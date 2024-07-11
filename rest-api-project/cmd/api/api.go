package api

import (
	"fmt"
	"net/http"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/db"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/services/hello"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
)

type Server struct {
	port string
	cfg  types.PostgresConfig
}

func NewServer(port string) *Server {
	cfg := &types.PostgresConfig{
		Username: "postgres",
		Password: "mysecretpassword",
		Host:     "localhost",
		Port:     5432,
		DB:       "dummy",
		Sslmode:  "disable",
	}

	_, err := db.Connect(*cfg)
	if err != nil {
		fmt.Println(err)
	}
	return &Server{port: port, cfg: *cfg}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	helloHandler := hello.NewHello()
	helloHandler.RegisterRoutes(router)

	return http.ListenAndServe(":"+s.port, router)
}
