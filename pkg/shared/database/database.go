package database

import (
	"fmt"
	"rest/config"
	"sync"

	"github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var database *Database
var dbConnOnce sync.Once

func GetDatabase() *Database {
	dbConnOnce.Do(func() {
		conn, err := dbr.Open("pgx", config.Database.GetUrl(), nil)
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}
		sess := conn.NewSession(nil)
		database = &Database{sess}
	})
	return database
}

type Database struct {
	*dbr.Session
}
