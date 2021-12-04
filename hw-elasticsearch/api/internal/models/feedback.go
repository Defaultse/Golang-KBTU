package models

import "time"

type Feedback struct {
	ID        int       `json:"id"`
	UserId    int       `json:"userId"`
	Author    int       `json:"author"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}
