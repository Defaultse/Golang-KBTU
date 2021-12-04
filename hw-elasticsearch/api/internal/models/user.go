package models

type User struct {
	ID          int    `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phonenumber" db:"phonenumber"`
}
