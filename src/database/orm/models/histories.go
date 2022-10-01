package models

import (
	"time"
)

type History struct {
	HistoryId     uint          `gorm:"type:uint;primaryKey;" json:"id,omitempty"`
	VehicleitemId uint          `gorm:"type:uint" json:"vehicleitem_id"`
	Vehicleitem   []Vehicleitem `gorm:"foreignKey:VehicleitemId;references:VehicleitemId;type:uint" json:"vehicleitem"`
	HistoryName   string        `json:"history_name,omitempty"`
	Favorite      bool          `json:"favorite"`
	RentFrom      string        `json:"rent_from,omitempty"`
	RentTo        string        `json:"rent_to,omitempty"`
	StatusReturn  string        `json:"status_return,omitempty"`
	Prepay        string        `json:"prepay,omitempty"`
	CreatedAt     time.Time     `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"default:now(); not null" json:"updated_at"`
}

type Histories []History
