package db

import (
	"database/sql"
	"log"

	// Import the SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase() *sql.DB {
	// Open or create the database file
	db, err := sql.Open("sqlite3", "internal/database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	// SQL statement to create the users table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	// Prepare the SQL statement for execution
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err) // Log any error that occurs during preparation
	}

	// Execute the prepared statement to create the table
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err) // Log any error that occurs during execution
	}
	return db
}
