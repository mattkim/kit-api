package models

import "time"

type ModelUUID struct {
	UUID      uint `gorm:"primary_key"`
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	ModelUUID
	Email string
}

// type SurveyItemAnswer struct {
// 	ModelUUID
// 	SurveyItemUUID    string
// 	CreatedByUser     User `gorm:"ForeignKey:CreatedByUserUUID"`
// 	CreatedByUserUUID string
// 	Answer            string
// }

// type SurveyItem struct {
// 	ModelUUID
// 	SurveyUUID string
// 	Question   string
// 	Answers    []SurveyItemAnswer `gorm:"ForeignKey:SurveyItemUUID"`
// }

// type Survey struct {
// 	ModelUUID
// 	SurveyItems []SurveyItem `gorm:"ForeignKey:SurveyUUID"`
// }

// type Location struct {
// 	ModelUUID
// 	Name             string
// 	Lat              float64
// 	Long             float64
// 	FormattedAddress string
// }

type EventDetail struct {
	ModelUUID
	Title       string
	Description string
	// Location     Location `gorm:"ForeignKey:LocationUUID"`
	// LocationUUID string
	StartTime time.Time
	EndTime   time.Time
}

type Event struct {
	ModelUUID
	CreatedByUser     User `gorm:"ForeignKey:CreatedByUserUUID"`
	CreatedByUserUUID string
	// Invitees          []User
	// Survey          Survey `gorm:"ForeignKey:SurveyUUID"`
	// SurveyUUID      string
	EventDetail     EventDetail `gorm:"ForeignKey:EventDetailUUID"`
	EventDetailUUID string
}
