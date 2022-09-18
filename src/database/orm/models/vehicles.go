package models

import "time"

type Vehicle struct {
	Vehicles_id   uint      `gorm:"primaryKey;autoIncrement" json:"vehicles_id"`
	Name          string    `json:"name"`
	Type_vehicles string    `json:"type_vehicles"`
	City          string    `json:"city"`
	Capacity      int       `json:"capacity"`
	Image         string    `json:"image"`
	Create_at     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Vehicles []Vehicle

type popular struct {
	VehiclesJoin int       `gorm:"column:vehicles_id"`
	vehicles     Histories `gorm:"foreignKey:VehicleJoin"`
}
