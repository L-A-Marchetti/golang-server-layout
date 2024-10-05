# Database Setup Process

## 1. Create `database.go` file in the `db` package

```go
package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func SetupDatabase() *sql.DB {
    // Open or create the database file
    db, err := sql.Open("sqlite3", "internal/database/forum.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create all tables
    createUsersTable(db)
    createPostsTable(db)
    createCommentsTable(db)
    createCategoriesTable(db)
    createPostCategoriesTable(db)
    createLikesDislikesTable(db)
    createImagesTable(db)

    return db
}
```
## 2. Create a function for each table
Example for the users table:

```go
// createUsersTable creates the users table in the database.
// It defines the schema for the users table, including fields for id, email, username, password, role, and creation timestamp.
func createUsersTable(db *sql.DB) {
    // SQL statement to create the users table if it does not already exist
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT, // Unique identifier for each user
        email TEXT NOT NULL UNIQUE,            // User's email address (must be unique)
        username TEXT NOT NULL UNIQUE,         // User's username (must be unique)
        password TEXT NOT NULL,                 // User's password (should be hashed in practice)
        role TEXT NOT NULL,                     // User's role (e.g., admin, moderator, user)
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP // Timestamp of when the user was created
    );`

    // Call executeSQL to run the SQL statement and create the table
    executeSQL(db, createTableSQL)
}
```
## 3. Create a utility function to execute SQL queries

```go
// executeSQL prepares and executes a given SQL statement.
// It logs any errors that occur during preparation or execution.
func executeSQL(db *sql.DB, sql string) {
    // Prepare the SQL statement for execution
    statement, err := db.Prepare(sql)
    if err != nil {
        log.Fatal(err) // Log and terminate if there is an error preparing the statement
    }

    // Execute the prepared statement
    _, err = statement.Exec()
    if err != nil {
        log.Fatal(err) // Log and terminate if there is an error executing the statement
    }
}
```
## 4. Repeat the process for each table
Example for the posts table:

```go
func createPostsTable(db *sql.DB) {
    createTableSQL := `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT NOT NULL,
        body TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );`

    executeSQL(db, createTableSQL)
}
```
## 5. Create functions for all other tables
Follow the same pattern for comments, categories, post_categories, likes_dislikes, and images tables.
## 6. Call SetupDatabase() in the functions made to threat datas from database like this:

```go
package db

// GetUsers retrieves all users from the database
func GetUsers() ([]User, error) {
	db := SetupDatabase() // We ask for the db connection.
	defer db.Close() // Then we close the connection after the function is executed.

	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
```