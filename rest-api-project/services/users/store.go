package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/types"
	"github.com/OlyMahmudMugdho/go-projects/rest-api-project/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func LoadSchema() ([]byte, error) {
	schema, err := os.ReadFile(utils.LoadSchemaLocation() + "userTable.sql")

	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (s *Store) CreateDatabase() {
	schema, _ := LoadSchema()
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
