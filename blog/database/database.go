package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) InitDB(url string) error {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	d.db = db

	return nil
}

func (d *Database) CloseDB() error {
	return d.db.Close()
}
