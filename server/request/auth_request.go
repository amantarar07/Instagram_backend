package request

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"11111111"`
}

type SendOtpRequest struct{

	PhoneNumber string `json:"phonenumber"`

}

type VerifyOtpRequest struct{
	// PhoneNumber string `json:"phonenumber"`
	Otp    string `json:"otp"`
}


