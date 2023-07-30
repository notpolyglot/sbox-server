package structs

import "time"

type User struct {
	Id        int       `json:"id"`
	ID64      string    `json:"id64"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Money int `json:"money"`
}
