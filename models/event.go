package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/libuuid"
)

type Event struct {
	ModelUUID
	CreatedByUser     User      `json:"created_by_user"`
	CreatedByUserUUID string    `gorm:"not null" json:"created_by_user_uuid"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Invitees          []User    `gorm:"many2many:events_invitees;" json:"invitees"`
	Messages          []Message `gorm:"ForeignKey:EventUUID" json:"messages"`
}

// GetEvent ...
func GetEvent(db *gorm.DB, uuid string) Event {
	event := Event{}
	db.Where("uuid = ?", uuid).Find(&event)
	// TODO: Is there a better way of doing this instead of using association?
	db.Model(&event).Association("Invitees").Find(&event.Invitees)
	db.Model(&event).Association("Messages").Find(&event.Messages)

	for i, _ := range event.Messages {
		db.Model(event.Messages[i]).Association("CreatedByUser").Find(&event.Messages[i].CreatedByUser)
	}

	db.Model(&event).Association("CreatedByUser").Find(&event.CreatedByUser)
	return event
}

// CreateEvent ...
func CreateEvent(db *gorm.DB, event Event) Event {
	t := time.Now()

	// Need to update these first.
	event.CreatedByUser = getUserIfExists(db, event.CreatedByUser)
	event.Invitees = getUsersIfExists(db, event.Invitees)
	event.Messages = createMessagesIfExists(db, event.Messages)

	// This is an update.
	if len(event.UUID) > 0 {
		event.UpdatedAt = t
		db.Save(&event)
		return event
	}

	event.UUID = libuuid.NewUUID()
	event.UpdatedAt = t
	event.CreatedAt = t

	db.Create(&event)

	return event
}
