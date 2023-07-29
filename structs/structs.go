package structs

import "time"

type LoginReq struct {
	Id string `json:"id64"`
}

type User struct {
	Id        int       `json:"id"`
	ID64      string    `json:"id64"`
	CreatedAt time.Time `json:"created_at"`
}
