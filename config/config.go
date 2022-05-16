package config

import (
	"os"
	"strings"
)

const (
	DEVELOPMENT = "dev"
	PRODUCTION  = "prod"
	TESTING     = "test"
)

var (
	AppName = Config("APP_NAME", "app")
	Mode    = ModeConf(Config("ENV", DEVELOPMENT))

	SecretKey = Config("SECRET_KEY", "secret")

	Server = ServerConf{
		Host:               Config("SERVER_HOST", "localhost"),
		Port:               Config("SERVER_PORT", "8080"),
		CorsAllowedOrigins: List(Config("CORS_ALLOWED_ORIGINS", "*")),
	}

	Database = DatabaseConf{
		Host:     Config("DB_HOST"),
		Port:     Config("DB_PORT"),
		User:     Config("DB_USER"),
		Password: Config("DB_PASSWORD"),
		Name:     Config("DB_NAME"),
		SSL:      Config("DB_SSL"),
	}

	Email = EmailConf{
		Host:     Config("EMAIL_HOST"),
		Port:     Config("EMAIL_PORT"),
		Username: Config("EMAIL_USERNAME"),
		Password: Config("EMAIL_PASSWORD"),
	}

	MigrationsDir = Config("MIGRATIONS_DIR", "migrations")
)

type ServerConf struct {
	Host               string
	Port               string
	CorsAllowedOrigins []string
}

type DatabaseConf struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSL      string
}

type EmailConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

type ModeConf string

func (m ModeConf) String() string {
	return string(m)
}

func (m ModeConf) IsDevelopment() bool {
	return m == DEVELOPMENT
}

func (m ModeConf) IsProduction() bool {
	return m == PRODUCTION
}

func (m ModeConf) IsTesting() bool {
	return m == TESTING
}

func (e ServerConf) GetUrl() string {
	return e.Host + ":" + e.Port
}

func (e DatabaseConf) GetUrl() string {
	return "host=" + e.Host + " user=" + e.User + " password=" + e.Password + " dbname=" + e.Name + " sslmode=" + e.SSL
}

func (e EmailConf) GetAddress() string {
	return e.Host + ":" + e.Port
}

func Config(key string, def ...string) string {
	env := os.Getenv(key)
	if env == "" && len(def) > 0 {
		env = def[0]
	}
	return env
}

func List(key string) []string {
	return strings.Split(Config(key), ",")
}
