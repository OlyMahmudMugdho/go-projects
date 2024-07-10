package main

import (
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/cmd/api"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/db"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
)

func main() {

	port := "8080"
	fmt.Println("server is running on port " + port)

	cfg := &types.PostgresConfig{
		Username: "postgres",
		Password: "mysecretpassword",
		Host:     "localhost",
		Port:     5432,
		DB:       "dummy",
		Sslmode:  "disable",
	}

	_, err := db.Connect(*cfg)

	server := api.NewServer(port)
	error := server.Run()

	if error != nil {
		log.Fatal(err)
	}
}
