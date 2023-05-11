package handler

import (
	"main/server/request"
	"main/server/services/refresh"
	"main/server/services/user"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func SearchHandler(context *gin.Context){


	utils.SetHeader(context)

	var username request.User
	utils.RequestDecoding(context,&username)

	
	user.Search(context,username)
}

func RefreshHandler(context *gin.Context){

	utils.SetHeader(context)

	refresh.RefreshService(context)

}


