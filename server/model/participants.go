package model

import "github.com/jinzhu/gorm"

type Participants struct {
	gorm.Model
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
	
}
