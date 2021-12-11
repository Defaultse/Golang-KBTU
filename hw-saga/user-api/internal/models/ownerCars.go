package models

type OwnerCars struct{
	UserID      int    `json:"user_id" db:"user_id"`
	CarId      int    `json:"car_id" db:"car_id"`
}