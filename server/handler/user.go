package handler

import (
	"fmt"
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
	fmt.Println("fullname: " ,userFullName)

	//call the service 
	user.UserFullNameService(context, userFullName)


}

func InstaUserNameHandler( context *gin.Context){

	utils.SetHeader(context)

	var InstaUserName request.InstaUserName
	utils.RequestDecoding(context, &InstaUserName)

	//if username already exists throw an error
	err := validation.CheckValidation(&InstaUserName)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	user.InstaUsernameService(context, InstaUserName)


}

func CreatePasswordHandler(context *gin.Context){

	utils.SetHeader(context)

	var userPassword  request.UserPassword
	utils.RequestDecoding(context, &userPassword)

		err := validation.CheckValidation(&userPassword)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	user.CreatePasswordService(context ,userPassword)


}

func UserLoginHandler(context *gin.Context){

	utils.SetHeader(context)

	var loginCred request.LoginCred
	utils.RequestDecoding(context, &loginCred)

	user.UserLoginService(context,loginCred)

	//give a jwt token to be used in headers for socket authentication

	



}
func SetUserBioHandler(context *gin.Context){

	utils.SetHeader(context)

	var bio request.Bio

	utils.RequestDecoding(context,&bio)

	user.SetUserBioService(context,bio)

}

func UploadPostHandler(context *gin.Context){

	utils.SetHeader(context)

	var caption request.Caption
	utils.RequestDecoding(context,&caption)

	user.UploadPostService(context,caption)

	
}

func GetUserPostsHandler(context *gin.Context){

	utils.SetHeader(context)
	
	user.GetUserPostService(context)
	

	
}

func LikePostHandler(context *gin.Context){

	utils.SetHeader(context)

	var like request.Like
	utils.RequestDecoding(context,&like)

	user.LikePostService(context,like)
}

func Comment_on_PostHandler(context *gin.Context){


	utils.SetHeader(context)

	var comment request.Comment
	utils.RequestDecoding(context,&comment)
	user.CommentOnPostService(context,comment)


}

func LikeCommentHandler(context *gin.Context){

	utils.SetHeader(context)

	var like request.Like
	utils.RequestDecoding(context,&like)

	user.LikeCommentService(context,like)
}

func FollwerUserHandler(context *gin.Context){


	utils.SetHeader(context)

	var otheruser request.User

	utils.RequestDecoding(context,&otheruser)

	user.FollowUserService(context,otheruser)
}