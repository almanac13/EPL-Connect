package models

import "time"

type ChatRoom struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdById"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"roomId"`
	SenderID  string    `json:"senderId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
