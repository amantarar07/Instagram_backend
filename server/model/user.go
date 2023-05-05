package model

type User struct {
	User_Id string  `json:"user_id" gorm:"default:uuid_generate_v4();unique"`
	Email    string `gorm:"type:varchar(250);UNIQUE"`
	Password string `gorm:"type:varchar(250);"`
	FullName string `gorm:"type:varchar(250);"`
	UserName string `gorm:"type:varchar(250);"`
	PhoneNumber string `json:"phonenumber"`
	Bio 	 string  `json:"bio"`
	IsActive string   `json:"is_active"`
}




