package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User => users
// OrderDetail => order_details
// Snake Case (lower_case jamak)
type User struct {
	ID        	uuid.UUID 		`gorm:"primary_key;column:id;type:char(36);<-:create"`
	Name      	string			`gorm:"column:name;size:255;not null;"`
	Email     	string			`gorm:"column:email;size:255;not null;unique"`
	Password  	string			`gorm:"column:password;size:255;not null;unique"`
	CreatedAt 	time.Time		`gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt 	time.Time		`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt	gorm.DeletedAt	`gorm:"column:deleted_at"`
}

// If we want to custom table Name
// func (u *User) TableName() string {
// 	return "users"
// }

// BeforeCreate hook to generate UUID before saving to the database
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}