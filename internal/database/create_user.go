package db

import (
	"fmt"
	"net/http"
)

// CreateUser handles the form submission to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		db := SetupDatabase()
		defer db.Close()

		// Start a new transaction
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "Error starting transaction", http.StatusInternalServerError)
			return
		}

		// Execute the insert statement directly within the transaction
		insertSQL := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
		_, err = tx.Exec(insertSQL, username, email, password)
		if err != nil {
			tx.Rollback() // Rollback if execution fails
			http.Error(w, "Error executing statement", http.StatusInternalServerError)
			return
		}

		// Commit the transaction if everything is successful
		if err = tx.Commit(); err != nil {
			http.Error(w, "Error committing transaction", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User %s created successfully!", username)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
