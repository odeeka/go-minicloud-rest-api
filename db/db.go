// Initialize the database
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "minicloud.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

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
