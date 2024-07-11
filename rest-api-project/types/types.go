package types

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
	Sslmode  string
}
