package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"

	handlerFn := http.HandlerFunc(SimpleApi)
	http.Handle("GET /", Logger(handlerFn))

	// chaining of middlewares
	http.Handle("GET /chaining", Logger(HeaderParser(handlerFn)))

	log.Printf("server is running on port %v \n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v \n", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func HeaderParser(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func SimpleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "hello world",
		},
	)
}
