// Initialize the database
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// DB is a globally accessible database handle
var DB *sql.DB

// InitDB opens a connection to the SQLite database and creates the required tables if they don't exist
func InitDB() {

	var err error

	// Open a connection to a SQLite database file
	DB, err = sql.Open("sqlite3", "minicloud.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	// Set the maximum number of open and idle connections
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create required tables
	createTables()
}

// createTables defines and creates the database tables
// for VMs, storage volumes, and user accounts
func createTables() {
	createVmsTable := `
	CREATE TABLE IF NOT EXISTS vms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		image TEXT NOT NULL,
		cpu REAL,
		memory INTEGER,
		ports TEXT,
		env TEXT,
		container_id TEXT
	);
	`

	_, err := DB.Exec(createVmsTable)

	if err != nil {
		panic("Could not create 'vms' table.")
	}

	// Create the 'storages' table to store virtual disk info
	createStoragesTable := `
	CREATE TABLE IF NOT EXISTS storages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		size_gb INTEGER NOT NULL,
		vm_id INTEGER,
		container_id TEXT
	);
	`

	_, err = DB.Exec(createStoragesTable)

	if err != nil {
		panic("Could not create 'storages' table.")
	}

	// Create the 'accounts' table for basic authentication
	createAccountsTable := `
CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);`

	_, err = DB.Exec(createAccountsTable)
	if err != nil {
		panic("Could not create 'accounts' table.")
	}

}
