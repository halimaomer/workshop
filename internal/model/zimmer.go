package model

type Zimmer struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"    json:"id"`
	Preis        float64 `gorm:"type:decimal(10,2);not null" json:"preis"`
	Zimmernummer string  `gorm:"not null"                    json:"zimmernummer"`
	HotelID      uint    `gorm:"column:hotel_id;index"       json:"hotel_id"`
}

func (Zimmer) TableName() string {
	return "hotel.zimmer"
}
