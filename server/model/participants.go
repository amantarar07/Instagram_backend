package model

import "github.com/jinzhu/gorm"

type Participants struct {
	gorm.Model
	RoomID int
	Room   Room `gorm:"references:RoomID"`
	UserID int
	User   User `gorm:"references:UserID"`
}
