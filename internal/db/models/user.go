package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"userId"`
	EmailId   string    `json:"emailId"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
}
