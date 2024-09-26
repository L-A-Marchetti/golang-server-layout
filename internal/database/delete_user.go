package db

import "net/http"

// DeleteUser handles the request to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")

		db := SetupDatabase()
		defer db.Close()

		// Start a new transaction
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "Error starting transaction", http.StatusInternalServerError)
			return
		}

		// Execute the delete statement directly within the transaction
		deleteSQL := `DELETE FROM users WHERE id = ?`
		_, err = tx.Exec(deleteSQL, id)
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

		// Redirect back to the index page after deletion
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
