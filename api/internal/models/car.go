package models

type (
	Car struct {
	ID      int    `json:"id"`
	OwnerId int    `json:"ownerId"`
	VehicleType    int   `json:"vehicleType"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Color   int    `json:"color"`
	Year           int     `json:"year"`
	EngineCapacity float32 `json:"engineCapacity"`
	Description    string  `json:"description"`
	Price          int     `json:"price"`
	Available      bool    `json:"available"`
	}

	CarsFilter struct {
		Query *string `json:"query"`
	}

	Brand struct {
		BrandId int `json:"brandId"`
		BrandName string `json:"brandName"`
	}

	Model struct {
		ModelId    int    `json:"modelId"`
		ModelName string `json:"modelName"`
		BrandId int `json:"brandId"`
	}

	VehicleType struct {
		TypeId   int    `json:"typeId"`
		TypeName string `json:"typeName"`
	}

	Color struct {
		ColorId int `json:"colorId"`
		ColorName string `json:"colorName"`
	}

)

type Type int

const (
	Sedan Type = iota
	SUV
	CUV
	Coupe
)
