package models

import "time"

type ModelUUID struct {
	UUID      string     `gorm:"primary_key" json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TODO: we need a way to generically apply this to all models.
// func Create(db *gorm.DB, model interface{}) {
// 	m := model.(ModelUUID)
// 	t := time.Now()
// 	m.UUID = libuuid.NewUUID()
// 	m.UpdatedAt = t
// 	m.CreatedAt = t
// 	db.Create(&m)
// }
