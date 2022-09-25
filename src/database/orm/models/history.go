package models

import "time"

type Histories struct {
	Histories_id   uint      `gorm:"primaryKey;autoIncrement" json:"histories_id"`
	User_id        string    `gorm:"type:varchar;not null" json:"user_id"`
	Vehicles_id    int       `gorm:"type:int;not null" json:"vehicles_id"`
	Date_order_at  string    `gorm:"type:date;not null" json:"date_order_at"`
	Date_order_end string    `gorm:"type:date;not null" json:"date_order_end"`
	Price          string    `gorm:"type:varchar;not null" json:"price"`
	Payment        string    `gorm:"type:varchar;not null" json:"payment"`
	Status         string    `gorm:"type:varchar;not null" json:"status"`
	Create_at      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type HistoriesAll []Histories
