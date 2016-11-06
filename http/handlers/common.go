package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorMessage struct {
	Message string
}

func writeJSON(w http.ResponseWriter, d interface{}) {
	fmt.Println("writeJSON")
	fmt.Printf("%+v\n", w.Header())
	writeDefaultHeaders(w)
	fmt.Printf("%+v\n", w.Header())
	b, err := json.Marshal(d)

	if err != nil {
		fmt.Fprintf(w, "%v", err)
	} else {
		fmt.Fprintf(w, "%s", b)
	}
	fmt.Printf("%+v\n", w.Header())
}

func writeDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
