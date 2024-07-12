package main

import (
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/cmd/api"
)

func main() {

	port := "8080"
	fmt.Println("server is running on port " + port)

	server := api.NewServer(port)
	error := server.Run()

	if error != nil {
		log.Fatal(error)
	}
}
