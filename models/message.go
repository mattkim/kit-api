package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/libuuid"
)

// Message ...
type Message struct {
	ModelUUID
	CreatedByUser     User   `json:"created_by_user"`
	CreatedByUserUUID string `gorm:"not null" json:"created_by_user_uuid"`
	EventUUID         string `gorm:"not null" json:"event_uuid"`
	Message           string `gorm:"not null" json:"message"`
}

// CreateMessage ...
func CreateMessage(db *gorm.DB, message Message) Message {
	t := time.Now()

	// This is an update.
	if len(message.UUID) > 0 {
		message.UpdatedAt = t
		db.Save(&message)
		return message
	}

	message.UUID = libuuid.NewUUID()
	message.UpdatedAt = t
	message.CreatedAt = t

	db.Create(&message)
	return message
}

// CreateMessages ...
func CreateMessages(db *gorm.DB, messages []Message) []Message {
	var res []Message

	for _, m := range messages {
		res = append(res, CreateMessage(db, m))
	}

	return res
}

func createMessagesIfExists(db *gorm.DB, messages []Message) []Message {
	if len(messages) == 0 {
		return messages
	}

	return CreateMessages(db, messages)
}
