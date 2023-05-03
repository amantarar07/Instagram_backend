package model


type UserAuthSessions struct {


	SessionId string `json:"sessionId" gorm:"default:uuid_generate_v4()"`
	PhoneNumber string `json:"phoneNumber"`
	VerifiedPhoneNumber string `json:"verifiedPhoneNumber"`
	Email 	string `json:"email"`
	VerifiedEmail string `json:"verifiedEmail"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Password string `json:"password"`

}