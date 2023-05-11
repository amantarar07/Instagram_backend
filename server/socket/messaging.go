package socket

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

func Messaging(s socketio.Conn,data map[string]string){

	var message model.Message
	//room_id from data 
	//user_id from token in header
	fmt.Println("messaging function called")

	roomId:=data["room_id"]
	token:=s.RemoteHeader().Get("authToken")
	claims,_:=utils.DecodeToken(token)

	
	

	message.Room_id=roomId
	message.Sender_id=claims.Id
	message.Text=data["text"]

	utils.SocketServerInstance.BroadcastToRoom("/",message.Room_id,"msg",message.Text)

	er:=db.CreateRecord(&message)
	if er!=nil{
		response.SocketResponse(
			"Failure",
			"server error",
			s,
		)
	}


}

func SendFileMessage(s socketio.Conn,room map[string]string){


	var file model.Files
	headerToken:=s.RemoteHeader().Get("auth")

	
	claims,_:=utils.DecodeToken(headerToken)
	// query:="select * from files where user_id='"+claims.Id+"' order by created_at LIMIT 1;"
	query:="select * from posts where user_id='"+claims.Id+"' order by created_at LIMIT 1;"



	db.QueryExecutor(query,&file)

	utils.SocketServerInstance.BroadcastToRoom("/",room["room_id"],"message",file.Path)
	


}        