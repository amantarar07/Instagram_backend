package handler

import (
	"main/server/request"
	"main/server/response"
	auth "main/server/services/authentication"
	"main/server/utils"
	"main/server/validation"

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

func UserSignupPhone(context *gin.Context){

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