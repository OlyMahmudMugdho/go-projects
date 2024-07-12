package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
	_ "github.com/lib/pq"
)

func ConnectionString(p types.PostgresConfig) string {
	return "user=" + p.Username + " " + "dbname=" + p.DB + " " + "sslmode=" + p.Sslmode + " " + "password=" + p.Password
}

func Connect(config types.PostgresConfig) (conn *sql.DB, err error) {
	fmt.Println(ConnectionString(config))
	db, err := sql.Open("postgres", ConnectionString(config))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("can't connect")
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}

	defer db.Close()
	return db, err
}
