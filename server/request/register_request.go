package request

type UserRegisterRequest struct {

	UserName string `json:"user_name" binding:"required"`
	FullName string `json:"full_name"`
	Email    string `json:"email" gorm:"omitempty"`
	PhoneNumber string `json:"phone_number" gorm:"omitempty"`
	Password string `json:"password" binding:"required" example:"11111111"`
	
}

type RegisterEmail struct{

	Email string `json:"email" binding:"required"`
}
type RegisterPhone struct{

	PhoneNumber string `json:"phonenumber"`
}
type FullName struct{

	FullName string `json:"full_name"`
}
type EmailOtp struct{

	Emailotp string `json:"email_otp"`
}

type InstaUserName struct{

	InstaUserName string `json:"insta_user_name"`
}
type UserPassword struct{

	UserPassword string `json:"user_password"`
}