package models

import (
	"time"
)

//Struct ke data sesuai dengan data di database

type Vehicle struct {
	VehicleId    uint      `gorm:"type:uint;primaryKey;" json:"vehicles_id"`
	Vehicle_Name string    `json:"vehicle_name,omitempty"`
	City         string    `json:"city,omitempty"`
	Status       string    `json:"status,omitempty"`
	Capacity     string    `json:"capacity,omitempty"`
	Type         string    `json:"type,omitempty"`
	Description  string    `json:"description,omitempty"`
	Price        int       `json:"price,omitempty"`
	Popular      int       `json:"popular,omitempty"`
	CreatedAt    time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Vehicleitem struct {
	VehicleitemId uint      `gorm:"type:uint;primaryKey;" json:"vehicles_id"`
	VehicleId     uint      `gorm:"type:uint" json:"vehicle_id"`
	Vehicle       []Vehicle `gorm:"foreignKey:VehicleId;references:VehicleId;type:uint" json:"vehicle"`
	UserId        string    `gorm:"type:uuid" json:"user_id"`
	User          User      `json:"user"`
	Amount        int       `json:"amount,omitempty"`
	Total         int       `json:"total,omitempty"`
	CreatedAt     time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Vehicles []Vehicle
type Vehicleitems []Vehicleitem
