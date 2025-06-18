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

	// createEventsTable := `
	// CREATE TABLE IF NOT EXISTS events (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT NOT NULL,
	// 	description TEXT NOT NULL,
	// 	location TEXT NOT NULL,
	// 	dateTime DATETIME NOT NULL,
	// 	user_id INTEGER,
	// 	FOREIGN KEY(user_id) REFERENCES users(id)
	// )
	// `

	// _, err = DB.Exec(createEventsTable)

}
