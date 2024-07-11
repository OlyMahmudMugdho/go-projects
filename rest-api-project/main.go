package main

import (
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/cmd/api"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/utils/dotenv"
)

func main() {

	dotenv.Load()

	vars := make(map[string]string)

	vars["foo"] = "bar"

	for key, value := range vars {
		fmt.Println(key + " : " + value)
	}

	port := "8080"
	fmt.Println("server is running on port " + port)

	server := api.NewServer(port)
	error := server.Run()

	if error != nil {
		log.Fatal(error)
	}
}
