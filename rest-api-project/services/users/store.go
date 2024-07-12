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
	return &Store{db: db}
}

func (s *Store) LoadSchema() ([]byte, error) {
	data, err := os.ReadFile("./services/users/userTable.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

func (s *Store) CreateDatabase() {
	schema, _ := s.LoadSchema()
	_, error := s.db.Exec(string(schema))

	if error != nil {
		fmt.Println("error occurred in store")
		return
	}

	fmt.Println("user table created")
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
