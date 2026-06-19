package model

type Standort struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"    json:"id"`
	Strasse    string `gorm:"not null"                    json:"strasse"`
	Hausnummer string `gorm:"not null"                    json:"hausnummer"`
	PLZ        string `gorm:"column:plz;not null;index"   json:"plz"`
	Ort        string `gorm:"not null"                    json:"ort"`
	Land       string `gorm:"not null"                    json:"land"`
	HotelID    uint   `gorm:"column:hotel_id;uniqueIndex" json:"hotel_id"`
}

func (Standort) TableName() string {
	return "hotel.standort"
}
