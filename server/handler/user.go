package handler

import (
	"main/server/request"
	"main/server/response"
	auth "main/server/services/authentication"
	"main/server/services/user"
	"main/server/utils"
	"main/server/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CreateUserHandler(context *gin.Context) {

// 	utils.SetHeader(context)

// 	var createUser request.CreateUserRequest

// 	utils.RequestDecoding(context, &createUser)

// 	err := validation.CheckValidation(&createUser)
// 	if err != nil {
// 		response.ErrorResponse(context, 400, err.Error())
// 		return
// 	}

// 	user.CreateUserService(context, createUser)
// }

func UserSignupPhoneHandler(context *gin.Context){

	utils.SetHeader(context)

	var userPhone request.RegisterPhone

	utils.RequestDecoding(context, &userPhone)

	err := validation.CheckValidation(&userPhone)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	auth.SendOtpService(context,request.SendOtpRequest(userPhone))
	//if otp gets verified only then create the user


}

func UserSignupEmailHandler(context *gin.Context){

	utils.SetHeader(context)
	
	var userEmail request.RegisterEmail
	utils.RequestDecoding(context, &userEmail)

	err := validation.CheckValidation(&userEmail)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	auth.SendEmailOtpService(context, userEmail.Email)

	//set another cookie for email address

	cookie:=&http.Cookie{Name: "UserEmail", Value:userEmail.Email}

	http.SetCookie(context.Writer,cookie)


}




func UserFullnameHandler(context *gin.Context){

	utils.SetHeader(context)

	var userFullName request.FullName

	utils.RequestDecoding(context, &userFullName)

	//call the service 
	user.UserFullNameService(context, userFullName)

	//set the cookie having value of phonenumber/email

}

func InstaUserNameHandler( context *gin.Context){


	
}