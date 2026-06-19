package handler

type CreateHotelRequest struct {
	Name     string          `json:"name" validate:"required,max=64"`
	Standort StandortRequest `json:"standort" validate:"required"`
	Zimmer   []ZimmerRequest `json:"zimmer"`
}
