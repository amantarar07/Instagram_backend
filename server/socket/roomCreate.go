package socket

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

func RoomCreate(s socketio.Conn,data map[string]string) {
	fmt.Println("inside room creation process...")
	// Get the user ID from the query params

	headerToken:=s.RemoteHeader().Get("authToken")
	claims,err:=utils.DecodeToken(headerToken)
	if err!=nil{

		fmt.Println("error in token decoding",err)
	}

	
	// s.Emit("hello","hi there")
	// utils.SocketServerInstance.BroadcastToNamespace("/","reply",data["message"])

	var room model.Room
	room.Creator=claims.Id
	room.Name=data["name"]
	db.CreateRecord(&room)

	s.Join(room.RoomID)
	s.Emit("createRoom",claims.Id+" :is connected to the room")



	//update the participant table with roomid and user id

	var participant model.Participants

	participant.RoomID=room.RoomID
	participant.UserID=claims.Id

	db.CreateRecord(&participant)

	// s.Emit("createRoom","room created successfully")
	response.SocketResponse("Success","room created Successfully",s)

	

	
	
}