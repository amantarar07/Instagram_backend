package socket

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

// "main/server/handler"


func ShowPastMessages(roomId string,event string,s socketio.Conn){


	fmt.Println("show past messages called...")
	var pastmessages []model.Message


	query:="select messages.sender_id,messages.text from messages where room_id='"+roomId+"' order by created_at;"
	er:=db.QueryExecutor(query,&pastmessages)
	if er!=nil{
		fmt.Println("db exe error")
	}
	fmt.Println("past messages of group",pastmessages)

	s.Emit("pastMessages",pastmessages)
	
	
}

func Typing(s socketio.Conn,data map[string]string){


	fmt.Println("typing func called")
	//get the user id from the token in header

	headerToken:=s.RemoteHeader().Get("authToken")

	claims,_:=utils.DecodeToken(headerToken)

	roomId:=data["room_id"]
	utils.SocketServerInstance.BroadcastToRoom("/",roomId,"typing",claims.Id+" is Typing")
	
	


}