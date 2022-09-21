package models

import "time"

type Histories struct {
	Histories_id   uint      `gorm:"primaryKey;autoIncrement" json:"histories_id"`
	User_id        string    `json:"user_id"`
	Vehicles_id    int       `json:"vehicles_id"`
	Date_order_at  string    `json:"date_order_at"`
	Date_order_end string    `json:"date_order_end"`
	Price          string    `json:"price"`
	Payment        string    `json:"payment"`
	Status         string    `json:"status"`
	Create_at      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type HistoriesAll []Histories
