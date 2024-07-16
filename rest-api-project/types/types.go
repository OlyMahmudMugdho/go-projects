package types

import (
	"time"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
	Sslmode  string
}

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStore interface {
	CreateUser(User) error
}
