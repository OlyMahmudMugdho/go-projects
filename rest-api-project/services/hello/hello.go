package hello

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Ok      *bool   `json:"ok"`
	Message *string `json:"message"`
}

func NewHello() *Hello {
	return &Hello{}
}

func NewHelloAllArgs(ok bool, message string) *Hello {
	return &Hello{
		Ok:      &ok,
		Message: &message,
	}
}

func (h *Hello) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", h.helloApi)
}

func (h *Hello) helloApi(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	hello := NewHelloAllArgs(true, "server is running")
	w.Header().Add("Content-Type", "application/json")
	encoder.Encode(&hello)
}
