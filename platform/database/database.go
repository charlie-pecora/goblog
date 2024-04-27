package database

import (
	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var Db *sqlx.DB

// connect db
func Connect(db_url string) error {
	dbConn, err := sqlx.Connect("postgres", db_url)
	if err != nil {
		return err
	}
	Db = dbConn
    return nil
}

