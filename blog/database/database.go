package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect(url string) error {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	d.db = db

	return nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) DB() *sql.DB {
	return d.db
}
