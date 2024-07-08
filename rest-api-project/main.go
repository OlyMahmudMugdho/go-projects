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

	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
