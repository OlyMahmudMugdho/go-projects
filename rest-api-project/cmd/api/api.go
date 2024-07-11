package api

import (
	"log"
	"net/http"
	"os"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/db"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/services/hello"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
	"github.com/joho/godotenv"
)

type Server struct {
	port string
	cfg  types.PostgresConfig
}

func NewServer(port string) *Server {
	envError := godotenv.Load()

	if envError != nil {
		log.Fatal(envError)
	}

	cfg := &types.PostgresConfig{
		Username: os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		DB:       os.Getenv("PG_DB"),
		Sslmode:  os.Getenv("PG_SSLMODE"),
	}

	_, err := db.Connect(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{port: port, cfg: *cfg}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	helloHandler := hello.NewHello()
	helloHandler.RegisterRoutes(router)

	return http.ListenAndServe(":"+s.port, router)
}
