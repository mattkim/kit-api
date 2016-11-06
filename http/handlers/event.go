package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/models"
)

// CreateEventRequest struct
type GetEventRequest struct {
	UUID string `json:"uuid"`
}

// EventHandler struct
type EventHandler struct {
	db *gorm.DB
}

// NewEventHandler constructs event handler
func NewEventHandler(db *gorm.DB) EventHandler {
	return EventHandler{db}
}

// CreateEventTestHandler handler
func (e EventHandler) CreateEventTestHandler(w http.ResponseWriter, r *http.Request) {
	event := models.CreateEvent(
		e.db,
		models.Event{
			Title:       "Test Title",
			Description: "Test Description",
			CreatedByUser: models.User{
				Email: "createdbyemail@gmail.com",
			},
			Invitees: []models.User{
				models.User{
					Email: "invitee@gmail.com",
				},
			},
		},
	)

	writeJSON(w, event)
}

// CreateEventHandler handler
func (e EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	v := models.Event{}
	json.NewDecoder(r.Body).Decode(&v)

	log.Printf("CreateEventHandler: %+v", v)

	// Required.
	if len(v.CreatedByUser.Email) == 0 {
		writeJSON(w, ErrorMessage{"Failed because email does not exist!"})
	}

	// Should be able to just pass it in
	event := models.CreateEvent(e.db, v)

	// Trying to refetch it here?
	event = models.GetEvent(e.db, event.UUID)
	log.Printf("CreateEvent: %+v", event)

	writeJSON(w, event)
}

// GetEventHandler ...
func (e EventHandler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")

	log.Printf("GetEventHandler: %+v", uuid)

	event := models.GetEvent(e.db, uuid)

	log.Printf("event: %+v", event)

	writeJSON(w, event)
}
