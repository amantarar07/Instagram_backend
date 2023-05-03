package server

import (
	_ "main/docs"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {
	
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	server.engine.POST("/user-signup-phone",handler.UserSignupPhoneHandler)
	server.engine.POST("/verify-phone",handler.VerifyPhoneOtpHandler)

	server.engine.POST("/user-signup-email",handler.UserSignupEmailHandler)
	server.engine.POST("/verify-email",handler.VerifyEmailOtpHandler)
	
	server.engine.POST("/enter-fullname",handler.UserFullnameHandler)

	server.engine.POST("/enter-user-name",handler.InstaUserNameHandler)

	server.engine.POST("/create-password",handler.CreatePasswordHandler)

	server.engine.POST("/user-login",handler.UserLoginHandler)
	

	
}
