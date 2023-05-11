package socket

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"main/server/utils"

	socketio "github.com/googollee/go-socket.io"
)

func RoomLeave(s socketio.Conn,data map[string]string){


	var participant model.Participants
	roomId := data["room_id"]
	s.Leave(roomId)

	token:=s.RemoteHeader().Get("authToken")
	claims,_:=utils.DecodeToken(token)
	
	
  //check if exists in participants table
	var ifExists bool
	query:="select exists(select * from participants where room_id='"+roomId+"' and user_id='"+claims.Id+"');"
	db.QueryExecutor(query,&ifExists)

	if ifExists{

		// delete the user from participants table
		db.InitDB().Where("user_id=? and room_id=?",claims.Id,roomId).Delete(&participant)
		// query:="delete * from participants where user_id='"+claims.Id+"' and room_id='"+roomId+"';"

		fmt.Println("deleted participant",participant)

		response.SocketResponse("Success","Successfully deleted from Room",s)
	}else{

		fmt.Println("not a member of this room")
		response.SocketResponse("Failed to delete participant","not a menmber of this room",s)
	}


}


