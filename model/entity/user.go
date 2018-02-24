package entity

import (
	"github.com/jinzhu/gorm"
)

// User example
type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name, omitempty"`
	LastName  string `gorm:"not null" json:"last_name, omitempty"`
	Email     string `json:"email, omitempty"`
	Login     string `gorm:"not null" json:"login, omitempty"`
	Password  string `gorm:"not null" json:"password, omitempty"`
}
