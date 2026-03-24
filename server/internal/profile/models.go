package profile

import "time"

type UserEntity struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Bio       string    `json:"bio"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
