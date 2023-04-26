package model

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	Name    string `json:"name"`
	Creator int
	User    User `gorm:"references:Creator"`
}
