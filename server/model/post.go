package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
	User    User `gorm:"references:UserID"`
}
