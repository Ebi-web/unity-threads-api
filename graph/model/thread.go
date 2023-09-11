package model

import "time"

type Thread struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	User      *User     `json:"user"`
	Posts     []*Post   `json:"posts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
