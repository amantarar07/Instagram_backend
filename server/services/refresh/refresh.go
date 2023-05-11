package refresh

import (
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func RefreshService(context *gin.Context){



	var newPosts []string

	cookie,_:=context.Request.Cookie("authToken")
	claims,_:=utils.DecodeToken(cookie.Value)

	// query:="select followers.follower_of from followers where follower ='"+claims.Id+"';"

	// db.QueryExecutor(query,&followers)

	// fmt.Println("followers:",followers)

	query:="select posts.post_id from posts where user_id IN (select followers.follower_of from followers where follower ='"+claims.Id+"') order by created_at;"

	db.QueryExecutor(query,&newPosts)
	fmt.Println("new posts",newPosts)

	response.ShowResponse("SUCCESS",200,"refersh success","",context)


}