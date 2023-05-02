package user

import (
	"main/server/model"
	"main/server/provider"
	"main/server/request"
	"net/http"

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


func UserFullNameService(context *gin.Context, fullName request.FullName){

	Phonecookie,err:=context.Request.Cookie("verifiedPhoneNumber")
	var Emailcookie *http.Cookie
	if err!=nil{
			//that means above cookie is not present
			 Emailcookie,_=context.Request.Cookie("verifiedEmail")
	}
	
	claims:=&model.Claims{Type: "fullname",FullName: fullName.FullName,PhoneNumber: Phonecookie.Value,Email: Emailcookie.Value}
	newtoken:=provider.GenerateToken(*claims, context)
	newcookie:=&http.Cookie{Name:"Fullname",Value:newtoken}

	//newcookie is set with phonenumber +fullname for further routes
	http.SetCookie(context.Writer,newcookie)


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