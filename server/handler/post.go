package handler

import (
	"main/server/request"
	"main/server/services/user"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func UploadPostHandler(context *gin.Context) {

	utils.SetHeader(context)

	var caption request.Caption
	utils.RequestDecoding(context, &caption)

	user.UploadPostService(context, caption)

}

func GetUserPostsHandler(context *gin.Context) {

	utils.SetHeader(context)

	user.GetUserPostService(context)

}

func LikePostHandler(context *gin.Context) {

	utils.SetHeader(context)

	var like request.Like
	utils.RequestDecoding(context, &like)

	user.LikePostService(context, like)
}

func Comment_on_PostHandler(context *gin.Context) {

	utils.SetHeader(context)

	var comment request.Comment
	utils.RequestDecoding(context, &comment)
	user.CommentOnPostService(context, comment)

}

func LikeCommentHandler(context *gin.Context) {

	utils.SetHeader(context)

	var like request.Like
	utils.RequestDecoding(context, &like)

	user.LikeCommentService(context, like)
}