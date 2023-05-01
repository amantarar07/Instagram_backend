package user

import (
	"main/server/request"

	"github.com/gin-gonic/gin"
)



func CreateUserService(context *gin.Context, decodedData request.RegisterPhone){

	// var user model.User
	// user.PhoneNumber=decodedData.PhoneNumber
	

	// if db.RecordExist("users", decodedData.PhoneNumber) {
	// 	response.ErrorResponse(context, 400, "User already exists")
	// 	return
	// }

	// err := db.CreateRecord(&user)
	// if err != nil {
	// 	response.ErrorResponse(context, 500, err.Error())
	// 	return
	// }
	// response.ShowResponse(
	// 	"Success",
	// 	200,
	// 	"User Created successfully",
	// 	user,
	// 	context,
	// )



	

}

// func SignUp_With_PhoneService(context *gin.Context , decodedData request.RegisterPhone){


// 	var user model.User
// 	user.PhoneNumber=decodedData.PhoneNumber
	

// 	if db.RecordExist("users", decodedData.PhoneNumber) {
// 		response.ErrorResponse(context, 400, "User already exists")
// 		return
// 	}

// 	err := db.CreateRecord(&user)
// 	if err != nil {
// 		response.ErrorResponse(context, 500, err.Error())
// 		return
// 	}
// 	response.ShowResponse(
// 		"Success",
// 		200,
// 		"User Created successfully",
// 		user,
// 		context,
// 	)

		

// }