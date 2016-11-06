package handlers

import "net/http"

// CORSHandler ...
func CORSHandler(w http.ResponseWriter, r *http.Request) {
	writeDefaultHeaders(w)
}
