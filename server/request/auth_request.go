package request

type LoginCred struct {

	Login_Input string `json:"login_input"`
	Password string `json:"password" binding:"required" example:"11111111"`
}

type SendOtpRequest struct{

	PhoneNumber string `json:"phonenumber"`

}

type VerifyOtpRequest struct{
	// PhoneNumber string `json:"phonenumber"`
	Otp    string `json:"otp"`
}





