package models

type User struct {
	ID          int    `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
	Role        Role   `json:"role" db:"role"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phoneNumber"`
}

type Role int

const (
	Owner Role = iota
	Client
)
