package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Sender_id string `json:"sender_id"`
	Room_id   string `json:"room_id"`
	Text  string `json:"text"`
}