package handler

import (
	"fmt"
	"main/server/socket"

	// "main/server/socket"

	socketio "github.com/googollee/go-socket.io"
)

// func SocketInit() *socketio.Server {
// 	server := socketio.NewServer(nil)
// 	SocketHandler(server)
// 	go server.Serve()
// 	return server
// }
func SocketHandler(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		
		//get the user_id from the token in the header

		return nil
	})

	 server.OnEvent("/","createRoom",socket.RoomCreate)
	 server.OnEvent("/","JoinRoom", socket.Roomjoin)
	 server.OnEvent("/","msg",socket.Messaging)
}


