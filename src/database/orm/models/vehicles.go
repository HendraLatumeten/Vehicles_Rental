package models

import "time"

type Vehicle struct {
	Vehicles_id   int       `gorm:"primaryKey;autoIncrement" json:"vehicles_id"`
	Name          string    `gorm:"type:varchar;not null" json:"name"`
	Type_vehicles string    `gorm:"type:varchar;not null" json:"type_vehicles"`
	City          string    `gorm:"type:varchar;not null" json:"city"`
	Capacity      int       `gorm:"type:int;not null" json:"capacity"`
	Image         string    `gorm:"type:varchar;not null" json:"image"`
	Orders        int       `gorm:"type:int" json:"orders"`
	Create_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Vehicles []Vehicle
