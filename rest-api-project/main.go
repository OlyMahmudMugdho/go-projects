package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func main() {
	router := http.NewServeMux()
	port := "8080"
	fmt.Println("server is running on port " + port)
	router.HandleFunc("GET /", helloApi)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func helloApi(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	hello := NewHello(true, "server is running")
	w.Header().Add("Content-Type", "application/json")
	encoder.Encode(&hello)
}
