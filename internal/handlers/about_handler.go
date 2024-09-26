package handlers

import (
	"html/template"
	"net/http"
)

// AboutHandler handles requests to the /about path.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the templates from files.
	tmpl, err := template.ParseFiles("web/pages/index.html", "web/templates/header.html", "web/templates/main.html", "web/templates/footer.html")
	if err != nil {
		http.Error(w, "Internal Server Error (Error parsing templates)", http.StatusInternalServerError)
		return
	}
	// Example of copying the index configuration with modifying th header title.
	aboutData := IndexData
	aboutData.Header.Title = "About"
	// Execute the template and pass any data needed (nil here for simplicity).
	err = tmpl.Execute(w, aboutData)
	if err != nil {
		http.Error(w, "Internal Server Error (Error executing template)", http.StatusInternalServerError)
		return
	}
}
