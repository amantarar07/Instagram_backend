package socket

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

func Roomjoin(s socketio.Conn,data map[string]string){

	var participant model.Participants
	roomId:=data["room_id"]
	if roomId == "" {
		response.SocketResponse(
			"Failure",
			"Room id not found",
			s,
		)
		return
	}
	//to do :-- check whether room exists or not  
	
	
	s.Join(roomId)
	
	response.SocketResponse(
		roomId,
		"User Successfully joined Room "+roomId,
		s,
	)
	utils.SocketServerInstance.BroadcastToRoom("/",roomId,"ack",data["user_id"]+"has joined the room")
	//add the client and user into participants table
	//check if already in that room

	var Exists bool
	query:="select exists(select * from participants where user_id ='"+ data["user_id"] + "'and room_id='"+data["room_id"]+"');"
	db.QueryExecutor(query,&Exists)
	fmt.Println("exists in room ",Exists)
	if Exists{

		//show the past messages in the room
		ShowPastMessages(roomId,"pastMessages",s)


		return
	}

	

	participant.UserID=data["user_id"]
	participant.RoomID=roomId
	


	er:=db.CreateRecord(&participant)
	if er!=nil{

		response.SocketResponse("server error",er.Error(),s)
	}

	


}






