package auth

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/provider"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var twilioClient *twilio.RestClient

func TwilioInit(password string) {
	twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: password,
	})
}

// func AdminRegisterService(context *gin.Context, adminRequest request.RegisterRequest) {

// 	var credential model.Credential
// 	credential.UserName = adminRequest.Username
// 	credential.Contact = adminRequest.Contact
// 	credential.Role = "admin"

// 	if db.RecordExist("credentials", adminRequest.Contact) {
// 		response.ErrorResponse(context, 400, "Admin already registerd")
// 		return
// 	}

// 	if db.RecordExist("users", adminRequest.Contact) {
// 		response.ErrorResponse(context, 400, "Admin cannot register as user")
// 		return
// 	}

// 	err := db.CreateRecord(&credential)
// 	if err != nil {
// 		response.ErrorResponse(context, 500, err.Error())
// 		return
// 	}

// 	response.Response(context, 200, credential)
// }

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
	PhoneClaims:=&model.Claims{Type: "phone_number",PhoneNumber: phoneNumber.PhoneNumber}

	PhoneToken:=provider.GenerateToken(*PhoneClaims,context)
	cookie:=&http.Cookie{Name: "phonenumber",Value: PhoneToken}
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
func VerifyOtpService(context *gin.Context, verifyOtp request.VerifyOtpRequest) {

//get the phone number from the token(inside cookie)

	cookie,_:=context.Request.Cookie("phonenumber")
	claims,_:=provider.DecodeToken(cookie.Value)

	fmt.Println("claims",claims)
	
	if CheckOtp("+91"+claims.PhoneNumber, verifyOtp.Otp) {
		fmt.Println("verification sucess")
		
		//redirect to fullname route



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

func UserFullNameService(context *gin.Context, fullName request.FullName){

	cookie,_:=context.Request.Cookie("phonenumber")
	claims,_:=provider.DecodeToken(cookie.Value)
	claims.FullName=fullName.FullName
	newtoken:=provider.GenerateToken(claims, context)
	newcookie:=&http.Cookie{Name:"Fullname",Value:newtoken}

	//newcookie is set with phonenumber +fullname for further routes
	http.SetCookie(context.Writer,newcookie)


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
