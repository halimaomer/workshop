package handler

type ZimmerRequest struct {
	Zimmernummer string  `json:"zimmernummer" validate:"required"`
	Preis        float64 `json:"preis" validate:"required,gt=0"`
}
