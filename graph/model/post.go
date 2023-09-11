package model

import "time"

type Post struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	User      *User     `json:"user"`
	Thread    *Thread   `json:"thread"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
