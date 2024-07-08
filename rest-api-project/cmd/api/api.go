package api

import (
	"encoding/json"
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
	router.HandleFunc("GET /", HelloApi)
	return http.ListenAndServe(":"+s.port, router)
}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	hello := hello.NewHello(true, "server is running")
	w.Header().Add("Content-Type", "application/json")
	encoder.Encode(&hello)
}
