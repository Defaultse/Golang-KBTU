package models

type (
	Car struct {
		ID             int     `json:"id" db:"id"`
		OwnerId        int     `json:"owner_id" db:"owner_id"`
		RentTypeId     int     `json:"rent_type_id" db:"rent_type_id"`
		CarTypeId      int     `json:"car_type_id" db:"car_type_id"`
		BrandId        int  `json:"brand_id" db:"brand_id"`
		ModelId        int  `json:"model_id" db:"model_id"`
		ColorId        int     `json:"color_id" db:"color_id"`
		Year           int     `json:"year" db:"year"`
		EngineCapacity float32 `json:"engine_capacity" db:"engine_capacity"`
		Description    string  `json:"description" db:"description"`
		Price          int     `json:"price" db:"price"`
	}

	FullCar struct {
		ID             int     `json:"id" db:"id"`
		OwnerId        int     `json:"owner_id" db:"owner_id"`
		RentType       string  `json:"rent_type" db:"rent_type"`
		CarType        string  `json:"car_type" db:"car_type"`
		Brand          string  `json:"brand" db:"brand"`
		Model          string  `json:"model" db:"model"`
		Color          string  `json:"color" db:"color"`
		Year           int     `json:"year" db:"year"`
		EngineCapacity float32 `json:"engine_capacity" db:"engine_capacity"`
		Description    string  `json:"description" db:"description"`
		Price          int     `json:"price" db:"price"`
	}
)
