package socket

import (
	"main/server/db"
	"main/server/model"
	"main/server/response"

	socketio "github.com/googollee/go-socket.io"
)

func Roomjoin(s socketio.Conn,data map[string]string){


	roomId:=data["roomId"]
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
	
	//add the client and user into participants table
	//check if already in that room

	var Exists bool
	query:="select exists(select * from participants where user_id ='"+ data["user_id"] + "'and room_id='"+data["room_id"]+"');"
	db.QueryExecutor(query,&Exists)
	if Exists{
		return
	}

	var participant model.Participants

	participant.UserID=data["user_id"]
	participant.RoomID=roomId


	db.CreateRecord(&participant)
}


// func Participants(s socketio.Conn,user_id string,room_id string){


// }