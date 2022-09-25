package models

import "time"

type User struct {
	User_id       string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"user_id"`
	Username      string    `gorm:"type:varchar;not null" json:"username"`
	Email         string    `gorm:"type:varchar;not null" json:"email"`
	Role          string    `gorm:"type:varchar(5);not null" json:"role"`
	Password      string    `gorm:"type:varchar;not null" json:"password,omitempty" validate:"required"`
	Gender        string    `gorm:"type:varchar(1);not null" json:"gender"`
	Address       string    `json:"address"`
	Phone_number  uint      `gorm:"type:int;not null" json:"phone_number"`
	Date_of_birth string    `gorm:"type:varchar;not null" json:"date_of_birth"`
	Image         string    `gorm:"type:varchar;not null" json:"image"`
	Create_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Users []User
