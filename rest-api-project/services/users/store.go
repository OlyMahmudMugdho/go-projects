package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	data, error := os.ReadFile("./userTable.sql")
	if error != nil {
		log.Fatal("unable to create database")
	}
	_, err := db.Exec(string(data))
	if err != nil {
		log.Fatal(err)
	}
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Query("INSERT INTO users (first_name, last_name, email) VALUES (?,?,?)", user.FirstName, user.LastName, user.Email)

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("user added to the database")
	return nil
}
