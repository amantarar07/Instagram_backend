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
