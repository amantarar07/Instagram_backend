package handler

import (
	"main/server/request"
	"main/server/response"
	auth "main/server/services/authentication"
	"main/server/utils"
	"main/server/validation"

	"github.com/gin-gonic/gin"
)

// func SendOtpHandler(context *gin.Context) {
// 	utils.SetHeader(context)
// 	var phoneNumber request.SendOtpRequest
// 	utils.RequestDecoding(context, &phoneNumber)
// 	fmt.Println("phoneNumber is", phoneNumber)
// 	auth.SendOtpService(context, phoneNumber)
// }

func VerifyPhoneOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var verifyRequest request.VerifyOtpRequest
	utils.RequestDecoding(context, &verifyRequest)

	auth.VerifyPhoneOtpService(context, verifyRequest)

}

func VerifyEmailOtpHandler(context *gin.Context){

	utils.SetHeader(context)

	var userOtp request.EmailOtp
	utils.RequestDecoding(context, &userOtp)

	err := validation.CheckValidation(&userOtp)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	auth.VerifyEmailOtpService(context, userOtp.Emailotp)

	


	//send the user to fullname endpoint


}