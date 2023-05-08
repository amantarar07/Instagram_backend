package model

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	Name    string `json:"name"`
	RoomID  string  `json:"room_id" gorm:"default:uuid_generate_v4();unique"`
	Creator string `json:"creater"`
	
}
