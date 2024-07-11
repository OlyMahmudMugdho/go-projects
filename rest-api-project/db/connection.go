package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
	_ "github.com/lib/pq"
)

func ConnectionString(p types.PostgresConfig) string {
	return "postgresql://" + p.Username + ":" + p.Password + "@" + p.Host + ":" + string(p.Port) + p.DB + "?=sslmode=" + p.Sslmode
}

func Connect(config types.PostgresConfig) (conn *sql.DB, err error) {
	db, err := sql.Open("postgres", ConnectionString(config))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connected")
	}
	return db, err
}
