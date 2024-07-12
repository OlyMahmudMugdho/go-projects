package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/db"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/services/hello"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/services/users"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
	"github.com/joho/godotenv"
)

type ApiServer struct {
	port string
	cfg  types.PostgresConfig
	db   *sql.DB
}

func NewServer(port string) *ApiServer {
	envError := godotenv.Load()

	if envError != nil {
		log.Fatal(envError)
	}

	var cfg = &types.PostgresConfig{
		Username: os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		DB:       os.Getenv("PG_DB"),
		Sslmode:  os.Getenv("PG_SSLMODE"),
	}

	var conn string = db.ConnectionString(*cfg)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	error := db.Ping()

	if error != nil {
		log.Fatal(error)
		fmt.Println("connection failed")
	} else {
		fmt.Println("connected to db")
	}

	return &ApiServer{port: port, cfg: *cfg, db: db}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	helloHandler := hello.NewHello()
	helloHandler.RegisterRoutes(router)

	userStore := users.NewStore(s.db)
	userStore.CreateDatabase()
	// userStore.ReadSchema()
	// userHandler := users.NewUserHandler(userStore)

	// userHandler.RegisterRoutes(router)

	return http.ListenAndServe(":"+s.port, router)
}
