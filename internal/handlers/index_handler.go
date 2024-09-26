package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler handles requests to the root of the server.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}
