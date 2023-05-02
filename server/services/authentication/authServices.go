package auth

import (
	"crypto/tls"
	"fmt"
	"main/server/db"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	gomail "gopkg.in/mail.v2"
)

var twilioClient *twilio.RestClient

func TwilioInit(password string) {
	twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: password,
	})
}



func SendOtpService(context *gin.Context, phoneNumber request.SendOtpRequest) {
	var exists1 bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE phone_number=?)"
	err := db.QueryExecutor(query, &exists1, phoneNumber.PhoneNumber)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	// Response
	if exists1 {
		response.ErrorResponse(context, 400, "Number already exists")
		return
	}
	ok, sid := sendOtp("+91" + phoneNumber.PhoneNumber)
	fmt.Println("SID is", sid)

	//set cookie of phone number token


	
	cookie:=&http.Cookie{Name: "phonenumber",Value: phoneNumber.PhoneNumber}
	http.SetCookie(context.Writer,cookie)
	fmt.Println("cookie is set",cookie)
	
	if ok {
		response.ShowResponse("Success", 200, "OTP send sucessfully", sid, context)
	}
}
func sendOtp(to string) (bool, *string) {
	fmt.Println("sahdvasasjfjasfjsaf")
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)

	params.SetChannel("sms")
	fmt.Println("service sid",os.Getenv("VERIFY_SERVICE_SID"))
	resp, err := twilioClient.VerifyV2.CreateVerification(utils.TWILIO_VERIFY_SERVICE_SID, params)
	fmt.Println("resp",resp)
	if err != nil {
		fmt.Println("bbkjfbkdsfbkaj")
		fmt.Println("err ",err)
		return false, nil
	} else {
		return true, resp.Sid
	}

}
func VerifyPhoneOtpService(context *gin.Context, verifyOtp request.VerifyOtpRequest) {

//get the phone number from the token(inside cookie)

	phonecookie,_:=context.Request.Cookie("phonenumber")
	
	if CheckOtp("+91"+phonecookie.Value,verifyOtp.Otp) {
		fmt.Println("verification sucess")
		
		
		verifiedPhoneCookie:=&http.Cookie{Name: "verifiedPhoneNumber",Value: phonecookie.Value}

		http.SetCookie(context.Writer,verifiedPhoneCookie)
		fmt.Println("verified cookie set")
		response.ShowResponse("Success",200,"phoneNumber Verified","",context)
		
		//set the cookie with verified number
		return

	} else {
		response.ErrorResponse(context, 401, "Verification Failed")
		return
	}
}

// OTP code verification
func CheckOtp(to string, code string) bool {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	resp, err := twilioClient.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		return false
	} else if *resp.Status == "approved" {
		return true
	} else {
		return false
	}
}


func SendEmailOtpService(context *gin.Context,toEmail string){

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "amantarar01@gmail.com")
  
	// Set E-Mail receivers
	m.SetHeader("To", toEmail)
  
	// Set E-Mail subject
	m.SetHeader("Subject", "Instagram Email verification")
  
	// Set E-Mail body. You can set plain text or html with text/html
	otp:=rand.Int()
	strOtp:= strconv.Itoa(otp)
	m.SetBody("text/plain",strOtp)
  
	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "amantarar01@gmail.com", "mdyrprmdvxpfxjjp")
  
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
	  fmt.Println(err)
	  panic(err)
	}

	//set a cookie with hash value of otp
	hash,_:=utils.HashPassword(strOtp)
	cookie:=&http.Cookie{Name:"otpHash", Value:hash}
	http.SetCookie(context.Writer,cookie)
  
	return

}

func VerifyEmailOtpService(context *gin.Context,otp string){

	//get the hash from the cookie value
	otpHashcookie,_:=context.Request.Cookie("otpHash")
	Emailcookie,_:=context.Request.Cookie("UserEmail")

	
	if(utils.CheckPasswordHash(otp,otpHashcookie.Value)){

		fmt.Println("email verified successfully")
		//set the cookie with verified email

		cookie:=&http.Cookie{Name:"verifiedEmail",Value:Emailcookie.Value}
		http.SetCookie(context.Writer,cookie)

		response.ShowResponse("Success",200,"Email verified successfully","",context)

		return


	}else{
		response.ShowResponse("Forbidden",403,"wrong otp","",context)
		return 
	}

	
}



// func LogoutService(context *gin.Context, tokenString string) {

// 	provider.DeleteCookie(context)
// 	var blacklist model.BlackListedToken
// 	blacklist.Token = tokenString
// 	db.CreateRecord(&blacklist)

// 	var user model.User
// 	claims, err := provider.DecodeToken(tokenString)
// 	if err != nil {
// 		response.ErrorResponse(context, 400, err.Error())
// 	}
// 	db.FindById(&user, &claims.RegisteredClaims.ID, "user_id")
// 	user.IsActive = false
// 	db.UpdateRecord(&user, &claims.RegisteredClaims.ID, "user_id")

// }
