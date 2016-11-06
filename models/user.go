package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pbk/kit-api/libuuid"
)

// User ...
type User struct {
	ModelUUID
	Email string `gorm:"not null;index;unique" json:"email"`
}

func CreateUser(db *gorm.DB, user User) User {
	t := time.Now()
	user.UUID = libuuid.NewUUID()
	user.UpdatedAt = t
	user.CreatedAt = t
	db.Create(&user)
	return user
}

func getUserIfExists(db *gorm.DB, user User) User {
	// TODO: validate user has required fields.
	res := User{}

	db.Where(user).First(&res)

	log.Printf("For input user: %+v, Found result: %+v", user, res)

	if len(res.UUID) == 0 {
		return CreateUser(db, user)
	}

	return res
}

func getUsersIfExists(db *gorm.DB, users []User) []User {
	var res []User

	for _, user := range users {
		res = append(res, getUserIfExists(db, user))
	}

	return res
}
