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
	store := &Store{db: db}
	store.CreateUserTable()
	return store
}

func LoadSchema() ([]byte, error) {
	schema, err := os.ReadFile(utils.LoadSchemaLocation() + "userTable.sql")

	if err != nil {
		return nil, err
	}
	return schema, nil
}

func (s *Store) CreateUserTable() {
	schema, _ := LoadSchema()
	_, error := s.db.Exec(string(schema))

	if error != nil {
		fmt.Println("error occurred in store")
		return
	}

	fmt.Println("user table created")
}

func (s *Store) CreateUser(user types.User) error {

	var query string = fmt.Sprintf(`INSERT INTO users (first_name, last_name, email) VALUES ('%s','%s','%s')`, user.FirstName, user.LastName, user.Email)

	_, err := s.db.Exec(query)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	fmt.Println("user added to the database")
	return nil
}
