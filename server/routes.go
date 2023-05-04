package server

import (
	_ "main/docs"
	"main/server/handler"
	"main/server/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {
	

	//aws connection
	sess := utils.ConnectAws()

	server.engine.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	   })


	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

//REGISTER/LOGIN  routes
	server.engine.POST("/user-signup-phone",handler.UserSignupPhoneHandler)

	server.engine.POST("/verify-phone",handler.VerifyPhoneOtpHandler)

	server.engine.POST("/user-signup-email",handler.UserSignupEmailHandler)

	server.engine.POST("/verify-email",handler.VerifyEmailOtpHandler)
	
	server.engine.POST("/enter-fullname",handler.UserFullnameHandler)

	server.engine.POST("/enter-user-name",handler.InstaUserNameHandler)

	server.engine.POST("/create-password",handler.CreatePasswordHandler)

	server.engine.POST("/user-login",handler.UserLoginHandler)


//POST related routes
	server.engine.POST("/upload-post",handler.UploadPostHandler)

	server.engine.GET("/get-posts",handler.GetUserPostsHandler)

	server.engine.POST("/like-post",handler.LikePostHandler)

	server.engine.POST("/comment-post",handler.Comment_on_PostHandler)
	

	
}
