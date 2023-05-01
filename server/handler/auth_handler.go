package handler

import (
	"main/server/request"
	auth "main/server/services/authentication"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// func SendOtpHandler(context *gin.Context) {
// 	utils.SetHeader(context)
// 	var phoneNumber request.SendOtpRequest
// 	utils.RequestDecoding(context, &phoneNumber)
// 	fmt.Println("phoneNumber is", phoneNumber)
// 	auth.SendOtpService(context, phoneNumber)
// }

func VerifyOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var verifyRequest request.VerifyOtpRequest
	utils.RequestDecoding(context, &verifyRequest)

	auth.VerifyOtpService(context, verifyRequest)

}