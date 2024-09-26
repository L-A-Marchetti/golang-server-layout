package handlers

import (
	"db"
	"html/template"
	"net/http"
)

// IndexHandler handles requests to the root of the server.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}
	// Parse the templates from files.
	tmpl, err := template.ParseFiles("web/pages/index.html", "web/templates/header.html", "web/templates/main.html", "web/templates/footer.html", "web/templates/users.html")
	if err != nil {
		http.Error(w, "Internal Server Error (Error parsing templates)", http.StatusInternalServerError)
		return
	}
	IndexData.Users = users
	// Execute the template and pass any data needed (nil here for simplicity).
	err = tmpl.Execute(w, IndexData)
	if err != nil {
		http.Error(w, "Internal Server Error (Error executing template)", http.StatusInternalServerError)
		return
	}
}
