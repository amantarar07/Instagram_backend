package user

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		}else{
			fmt.Println("login failed")
			response.ShowResponse("Forbidden", 403,"Wrong credentials","",context)
			return

		}


	}
	fmt.Println("userid ",user.User_Id)
	var claims model.Claims
	claims.Id=user.User_Id

	token :=utils.GenerateToken(claims)

	cookie:=&http.Cookie{Name:"authToken",Value:token}
	http.SetCookie(context.Writer,cookie)

	response.ShowResponse("success", 200,"login successful","",context)
}




func SetUserBioService(context *gin.Context,bio request.Bio){


		//extract the userid from the auth token

		cookie,_:=context.Request.Cookie("authToken")
		claims,_:=utils.DecodeToken(cookie.Value)

		var user model.User

		er:=db.FindById(&user,claims.ID,"user_id")
	 	if er!=nil{
			response.ShowResponse("server error", 500,er.Error(),"",context)
			return
		}
		user.Bio=bio.Bio

		//update the entry into the table

		db.UpdateRecord(&user,claims.ID,"user_id")

}

func GetUserBioService(context *gin.Context){

		//extract the userid from the auth token

	cookie,_:=context.Request.Cookie("authToken")
	claims,_:=utils.DecodeToken(cookie.Value)

	var user model.User

	er:=db.FindById(&user,claims.ID,"user_id")
	if er!=nil{

		response.ShowResponse("server error", 500,er.Error(),"",context)
		return
	}
	response.ShowResponse("success",200,"Bio of user fetched successfully",user.Bio,context)


}





func UploadPostService(c *gin.Context ,caption request.Caption){


	
	fmt.Println("upload image called")
    sess := c.MustGet("sess").(*session.Session)
    uploader := s3manager.NewUploader(sess)
    MyBucket := os.Getenv("BUCKET_NAME")
    fmt.Println("bucket",MyBucket)
    file, header, err := c.Request.FormFile("file")
    fmt.Println("file",file)
    filename := header.Filename//upload to the s3 bucket
    fmt.Println("filename",filename)

    up, err := uploader.Upload(&s3manager.UploadInput{
     Bucket: aws.String(MyBucket),
     ACL:    aws.String("public-read"),
     Key:    aws.String(filename),
     Body:   file,
    })
    fmt.Println("error",err)

    if err != nil {
     c.JSON(http.StatusInternalServerError, gin.H{
      "error":    "Failed to upload file",
      "uploader": up,
     })
     return
    }
    filepath := "https://" + MyBucket + "." + "s3-" + utils.MyRegion + ".amazonaws.com/" + filename
    c.JSON(http.StatusOK, gin.H{
     "filepath":    filepath,
    })

	//add the post entry inside post table

	var post model.Post
	post.Title=filename
	post.Path=filepath
	post.Caption=caption.CaptionText

	//get the user id from the token 
	
	Cookie,_:=c.Request.Cookie("authToken")

	fmt.Println("cookie",Cookie.Value)

	claims,_:=utils.DecodeToken(Cookie.Value)

	fmt.Println("claims",claims)
	post.UserID=claims.ID
	er:=db.CreateRecord(&post)
	if er!=nil{

		response.ShowResponse("server error",500,er.Error(),"",c)
	}

}

func GetUserPostService(context *gin.Context){

	//extract the userid from the auth token

	cookie,_:=context.Request.Cookie("authToken")
	claims,_:=utils.DecodeToken(cookie.Value)

	query:="select * from posts where user_id='"+claims.ID+"';"
	var userPosts []model.Post
	db.QueryExecutor(query,&userPosts)

	response.ShowResponse("success",200,"posts fetched successfully",userPosts,context)

	
}

func LikePostService(context *gin.Context,Like request.Like){


	//get the user id of the user who like
	cookie,_:=context.Request.Cookie("authToken")

	claims,_:=utils.DecodeToken(cookie.Value)

	var like model.Like
	like.PostID=Like.PostID
	like.User_Id=claims.ID
	er:=db.CreateRecord(&like)
	if er!=nil{
		response.ShowResponse("server error",500,er.Error(),"",context)
		return
	}

	

	//update the like count of the post
	var post model.Post
	db.FindById(&post,Like.PostID,"post_id")
	post.Likes+=1
	db.UpdateRecord(&post,Like.PostID,"post_id")
	response.ShowResponse("success",200,"post liked successfully","",context)


}

func CommentOnPostService(context *gin.Context,comment request.Comment){

	//get the user id of the user who like
	cookie,_:=context.Request.Cookie("authToken")

	claims,_:=utils.DecodeToken(cookie.Value)

	var comnt model.Comment
	comnt.CommentText=comment.Comment
	comnt.UserID=claims.ID
	comnt.PostID=comment.PostID
	//create comment entry in db
	err:=db.CreateRecord(&comnt)
	if err!=nil{
		response.ShowResponse("server error",500,err.Error(),"",context)
		return
	}

	//update the comment count in the post table

	var post model.Post
	er:=db.FindById(&post,comment.PostID,"post_id")
	if er!=nil{
		response.ShowResponse("server error",500,er.Error(),"",context)
	}
	post.Comments+=1

	db.UpdateRecord(&post,comment.PostID,"post_id")


	// query:="UPDATE posts SET comment=ARRAY_APPEND(comment,'"+comment.Comment+"') WHERE post_id='"+comment.PostID+"';"
	// er:=db.QueryExecutor(query,&commentedPost)
	// if er!=nil{
		
	// 	response.ShowResponse("server error",500,er.Error(),"",context)
	// 	return
	// }

	response.ShowResponse("success",200,"comment added successfully","",context)

	// commentOnPost.



}


func LikeCommentService(context *gin.Context,like request.Like){


		//get the user id of the user who like
		cookie,_:=context.Request.Cookie("authToken")

		claims,_:=utils.DecodeToken(cookie.Value)

		//create a like entry 

		var like_on_comment model.Like
		like_on_comment.User_Id=claims.ID
		like_on_comment.CommentID=like.CommentID

	
		er:=db.CreateRecord(&like_on_comment)
		if er!=nil{
			response.ShowResponse("server error",500,er.Error(),"",context)
			return
		} 

		var comnt model.Comment
		err:=db.FindById(&comnt,like_on_comment.CommentID,"comment_id")
		if err!=nil{
			response.ShowResponse("server error",500,err.Error(),"",context)
			return
		} 
		comnt.LikesCount+=1
		err1:=db.UpdateRecord(&comnt,like_on_comment.CommentID,"comment_id").Error
		if err1!=nil{

			response.ShowResponse("server error",500,err1.Error(),"",context)
		}

		response.ShowResponse("success",200,"like on comment added successfully","",context)
}

func FollowUserService(context *gin.Context,Otheruser request.User){

		//get the user id of the user who like
		cookie,_:=context.Request.Cookie("authToken")

		claims,_:=utils.DecodeToken(cookie.Value)

		var follower model.Followers

		follower.User_id=claims.ID
		follower.FollowerOf=Otheruser.User_id

		er:=db.CreateRecord(&follower)
	   if er!=nil{

		response.ShowResponse("server error",500,er.Error(),"",context)
		return 
	   }

	   

}
