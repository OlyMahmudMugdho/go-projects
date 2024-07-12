package users

import (
	"os"
	"testing"
)

func TestSchemaValidation(t *testing.T) {
	loadedSchema, err := os.ReadFile("./userTable.sql")

	var expectedSchema string = `CREATE TABLE IF NOT EXISTS USERS (
    ID SERIAL PRIMARY KEY,
    FIRST_NAME VARCHAR(255),
    LAST_NAME VARCHAR(255),
    EMAIL TEXT UNIQUE NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT NOW()
);`

	if err != nil {
		t.Errorf("can't load schema")
	}

	if string(loadedSchema) != expectedSchema {
		t.Errorf("invalid schema")
	}
}
