package hello

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewHello(ok bool, message string) *Hello {
	return &Hello{
		Ok:      ok,
		Message: message,
	}
}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	hello := NewHello(true, "server is running")
	w.Header().Add("Content-Type", "application/json")
	encoder.Encode(&hello)
}
