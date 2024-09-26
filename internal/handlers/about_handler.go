package handlers

import (
	"fmt"
	"net/http"
)

// AboutHandler handles requests to the /about path.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}
