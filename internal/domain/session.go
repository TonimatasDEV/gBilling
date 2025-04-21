package domain

import "time"

type Session struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	Exp       time.Time `json:"exp"`
}
