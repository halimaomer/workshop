package model

import "time"

type Hotel struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Version      int       `gorm:"default:0"               json:"version"`
	Name         string    `gorm:"not null;index"          json:"name"`
	Erzeugt      time.Time `gorm:"column:erzeugt"          json:"erzeugt"`
	Aktualisiert time.Time `gorm:"column:aktualisiert"     json:"aktualisiert"`
	Standort     *Standort `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE" json:"standort,omitempty"`
	Zimmer       []Zimmer  `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE" json:"zimmer,omitempty"`
}

func (Hotel) TableName() string {
	return "hotel.hotel"
}
