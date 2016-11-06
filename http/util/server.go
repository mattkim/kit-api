package util

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/http/handlers"
)

// NewServer ...
func NewServer(db *gorm.DB) {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/message/create", handlers.CORSHandler).Methods("OPTIONS")
	r.HandleFunc("/event/createtest", handlers.CORSHandler).Methods("OPTIONS")
	r.HandleFunc("/event/create", handlers.CORSHandler).Methods("OPTIONS")
	r.HandleFunc("/event", handlers.CORSHandler).Methods("OPTIONS")

	h := handlers.NewEventHandler(db)
	m := handlers.NewMessageHandler(db)

	r.HandleFunc("/message/create", m.CreateMessageHandler).Methods("POST")
	r.HandleFunc("/event/createtest", h.CreateEventTestHandler).Methods("POST")
	r.HandleFunc("/event/create", h.CreateEventHandler).Methods("POST")
	r.HandleFunc("/event", h.GetEventHandler).Methods("GET")

	log.Println("Running HTTP server on " + port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
