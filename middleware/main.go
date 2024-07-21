package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	http.HandleFunc("GET /", SimpleApi)
	log.Printf("server is running on port %v \n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func SimpleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "hello world",
		},
	)
}
