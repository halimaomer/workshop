package handler

type StandortRequest struct {
	Strasse    string `json:"strasse" validate:"required"`
	Hausnummer string `json:"hausnummer" validate:"required"`
	PLZ        string `json:"plz" validate:"required,len=5"`
	Ort        string `json:"ort" validate:"required"`
	Land       string `json:"land" validate:"required"`
}

type ZimmerRequest struct {
	Zimmernummer string  `json:"zimmernummer" validate:"required"`
	Preis        float64 `json:"preis" validate:"required,gt=0"`
}

type CreateHotelRequest struct {
	Name     string          `json:"name" validate:"required,max=64"`
	Standort StandortRequest `json:"standort" validate:"required"`
	Zimmer   []ZimmerRequest `json:"zimmer"`
}
