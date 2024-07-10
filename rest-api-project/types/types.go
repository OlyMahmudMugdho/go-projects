package types

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DB       string
	Sslmode  string
}
