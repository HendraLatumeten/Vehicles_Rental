package models

import "time"

type User struct {
	User_id       string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"user_id"`
	Username      string
	Email         string
	Password      string    `json:"password"`
	Gender        string    `json:"gender"`
	Address       string    `json:"address"`
	Phone_number  int64     `json:"phone_number"`
	Date_of_birth string    `json:"date_of_birth"`
	Image         string    `json:"image"`
	Create_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Users []User
