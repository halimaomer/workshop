package handler

type StandortRequest struct {
	Strasse    string `json:"strasse" validate:"required"`
	Hausnummer string `json:"hausnummer" validate:"required"`
	PLZ        string `json:"plz" validate:"required,len=5"`
	Ort        string `json:"ort" validate:"required"`
	Land       string `json:"land" validate:"required"`
}
