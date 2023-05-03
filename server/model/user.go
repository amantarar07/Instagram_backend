package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	User_Id string  `json:"user_id" gorm:"default:uuid_generate_v4()"`
	Email    string `gorm:"type:varchar(250);UNIQUE"`
	Password string `gorm:"type:varchar(250);"`
	FullName string `gorm:"type:varchar(250);"`
	UserName string `gorm:"type:varchar(250);"`
	PhoneNumber string `json:"phonenumber"`
	IsActive string   `json:"is_active"`
}


