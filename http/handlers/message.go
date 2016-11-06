package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/models"
)

// MessageHandler struct
type MessageHandler struct {
	db *gorm.DB
}

// NewMessageHandler constructs event handler
func NewMessageHandler(db *gorm.DB) MessageHandler {
	return MessageHandler{db}
}

// CreateEventHandler handler
func (e MessageHandler) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	v := models.Message{}
	json.NewDecoder(r.Body).Decode(&v)

	log.Printf(": %+v", v)

	if len(v.CreatedByUserUUID) == 0 || len(v.Message) == 0 {
		writeJSON(w, ErrorMessage{"Failed because user or message does not exist!"})
	}

	// // Should be able to just pass it in
	message := models.CreateMessage(e.db, v)
	log.Printf("CreateMessage: %+v", message)

	writeJSON(w, message)
}
