package user

import (
	"fmt"
	"io/ioutil"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)





func UserFullNameService(context *gin.Context, fullName request.FullName){


	fmt.Println("userfullnameservice called 0")

	sessionIdCookie,_:=context.Request.Cookie("sessionID")
	var usersession model.UserAuthSessions
	er:=db.FindById(&usersession,sessionIdCookie.Value,"session_id")
	if er!=nil{
		response.ShowResponse("server error",500,er.Error(),"",context)
	}
	fmt.Println("userfullnameservice called 1",fullName.FullName)

	usersession.FullName=fullName.FullName
	err:=db.UpdateRecord(&usersession,sessionIdCookie.Value,"session_id").Error
	if err!=nil{
		response.ShowResponse("server error",500,err.Error(),"",context)
	}	

	fmt.Println("fullname set successfully")
	response.ShowResponse("success", 200,"fullname set successfully","",context)


}

func InstaUsernameService(context *gin.Context,InstaUsername request.InstaUserName){


	sessionIdCookie,_:=context.Request.Cookie("sessionID")
	var usersession model.UserAuthSessions
	er:=db.FindById(&usersession,sessionIdCookie.Value,"session_id")
	if er!=nil{
		response.ShowResponse("server error",500,er.Error(),"",context)
	}
	usersession.Username=InstaUsername.InstaUserName
	err:=db.UpdateRecord(&usersession,sessionIdCookie.Value,"session_id").Error
	if err!=nil{
		response.ShowResponse("server error",500,err.Error(),"",context)
	}
	fmt.Println("Instausername set successfully")
	response.ShowResponse("success", 200,"Instausername set successfully","",context)


}

func CreatePasswordService(context *gin.Context,userPassword request.UserPassword){

	sessionIdCookie,_:=context.Request.Cookie("sessionID")
	var usersession model.UserAuthSessions
	er:=db.FindById(&usersession,sessionIdCookie.Value,"session_id")
	if er!=nil{
		response.ShowResponse("server error",500,er.Error(),"",context)
	}
	//hash the user password and then store in into session object
	passwordhash,_:=utils.HashPassword(userPassword.UserPassword)
	usersession.Password=passwordhash
	err:=db.UpdateRecord(&usersession,sessionIdCookie.Value,"session_id").Error
	if err!=nil{
		response.ShowResponse("server error",500,err.Error(),"",context)
	}
	fmt.Println("Password set successfully")


	//create user 
	var user model.User
	user.Email=usersession.VerifiedEmail
	user.PhoneNumber=usersession.VerifiedPhoneNumber
	user.FullName=usersession.FullName
	user.UserName=usersession.Username
	user.Password=usersession.Password
	er1:=db.CreateRecord(&user)
	if er1!=nil{
		response.ShowResponse("server error",500,er1.Error(),"",context)
	}
	response.ShowResponse("success", 200,"Password set successfully","",context)


}

func UserLoginService(context *gin.Context,credential request.LoginCred){



	var user model.User
	if utils.IsEmail(credential.Login_Input ){

		db.FindById(&user,credential.Login_Input,"email")
		if (utils.CheckPasswordHash(credential.Password,user.Password)){
			fmt.Println("login successful")

			//make this user active in the user table
			user.IsActive="true"
			err:=db.UpdateRecord(user,credential.Login_Input,"phone_number").Error
			if err!=nil{

				response.ShowResponse("server error", 500,err.Error(),"",context)
			}
			//give this user auth token
			var claims model.Claims
			claims.ID=user.User_Id

			token :=utils.GenerateToken(claims)

			cookie:=&http.Cookie{Name:"authToken",Value:token}
			http.SetCookie(context.Writer,cookie)

			response.ShowResponse("success", 200,"login successful","",context)
		}else{
			fmt.Println("login failed")
			response.ShowResponse("Forbidden", 403,"Wrong credentials","",context)

		}

	}else{


		
		db.FindById(&user,credential.Login_Input,"phone_number")

		if (utils.CheckPasswordHash(credential.Password,user.Password)){
			fmt.Println("login successful")
			//make this user active in the user table
			user.IsActive="true"
			err:=db.UpdateRecord(user,credential.Login_Input,"phone_number").Error
			if err!=nil{

				response.ShowResponse("server error", 500,err.Error(),"",context)
			}
			//give user a token
			var claims model.Claims
			claims.ID=user.User_Id

			token :=utils.GenerateToken(claims)

			cookie:=&http.Cookie{Name:"authToken",Value:token}
			http.SetCookie(context.Writer,cookie)
			
			response.ShowResponse("success", 200,"login successful","",context)
			return
		}else{
			fmt.Println("login failed")
			response.ShowResponse("Forbidden", 403,"Wrong credentials","",context)
			return

		}


	}


}

func UploadPostService(context *gin.Context ,filepath request.Filepath){


	fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    context.Request.ParseMultipartForm(1 << 50)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := context.Request.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()



}