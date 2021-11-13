package models

type Car struct {
	ID             int     `json:"id"`
	OwnerId        int     `json:"ownerId"`
	Type           Type    `json:"vehicleType"`
	Brand          string  `json:"brand"`
	Model          string  `json:"model"`
	Year           int     `json:"year"`
	EngineCapacity float32 `json:"engineCapacity"`
	Description    string  `json:"description"`
	Price          int     `json:"price"`
	Available      bool    `json:"available"`
}

type Type int

const (
	Sedan Type = iota
	SUV
	CUV
	Coupe
)
