// Package database provides a database connection
//
// This file contains the Database struct and the InitDB method.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Import pgx driver for database/sql package
)

// Database is a wrapper around the database connection.
// It contains a pointer to a sql.DB.
type Database struct {
	// The sql.DB is the database connection.
	// For more information on the sql.DB, see:
	// https://golang.org/pkg/database/sql/#DB
	//
	// We use an underscore to import the pgx driver.
	// We don't need to use the pgx driver directly.
	// We only need to import it so that the database/sql package can use it.
	// For more information on importing packages, see:
	// https://golang.org/doc/effective_go.html#blank_import
	db *sql.DB
}

// NewDatabase creates a new database.
// It returns a pointer to a Database.
func NewDatabase() *Database {
	return &Database{}
}

// InitDB initializes the database.
// It takes a database URL as an argument.
// It returns an error if the database fails to initialize.
//
// The database URL is a string that contains the database driver, the database host, the database port, the database name, the database user, and the database password.
// For example, "postgres://postgres:postgres@localhost:5432/blog?sslmode=disable".
func (d *Database) InitDB(url string) error {
	// sql.Open creates a new database connection.
	// It takes a database driver and a database URL as arguments.
	// It returns a pointer to a sql.DB and an error.
	// If the database connection fails to initialize, we return the error.
	//
	// For more information on sql.Open, see:
	// https://golang.org/pkg/database/sql/#Open
	db, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// d.db is a pointer to a sql.DB.
	// We set it to the pointer to the sql.DB that was returned by sql.Open.
	//
	// For more information on sql.DB, see:
	// https://golang.org/pkg/database/sql/#DB
	d.db = db

	return nil
}

// CloseDB closes the database connection.
// It returns an error if the database connection fails to close.
// It is good practice to close the database connection when you are done using it.
func (d *Database) CloseDB() error {
	return d.db.Close()
}

// GetDB returns a pointer to the database connection.
// It returns a pointer to a sql.DB.
// We use this method to get a pointer to the database connection so that we can pass it to the models.
func (d *Database) GetDB() *sql.DB {
	return d.db
}
