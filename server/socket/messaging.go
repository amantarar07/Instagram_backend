package socket

import (
	"fmt"
	"main/server/model"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

func Messaging(s socketio.Conn,data map[string]string){


	//room_id from data 
	//user_id from token in header
	fmt.Println("messaging function called")

	roomId:=data["room_id"]
	token:=s.RemoteHeader().Get("authToken")
	claims,_:=utils.DecodeToken(token)

	
	var message model.Message

	message.Room_id=roomId
	message.Sender_id=claims.Id

	utils.SocketServerInstance.BroadcastToRoom("/",data["room_id"],"msg",data["text"])
	


}