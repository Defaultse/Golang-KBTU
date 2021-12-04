package models

type (
	Car struct {
	ID      int    `json:"id" db:"id"`
	OwnerId int    `json:"owner_id" db:"owner_id"`
	RentTypeId    int   `json:"rent_type_id" db:"rent_type_id"`
	CarTypeId    int   `json:"car_type_id" db:"car_type_id"`
	ModelId   string `json:"model_id" db:"model_id"`
	ColorId   int    `json:"color_id" db:"color_id"`
	Year           int     `json:"year" db:"year"`
	EngineCapacity float32 `json:"engine_capacity" db:"engine_capacity"`
	Description    string  `json:"description" db:"description"`
	Price          int     `json:"price" db:"price"`
	}

	RentTypes struct {
		ID      int    `json:"id" db:"id"`
		RentTypes string `json:"rent_type" db:"rent_type"`
	}

	CarsFilter struct {
		Query *string `json:"query"`
	}

)

