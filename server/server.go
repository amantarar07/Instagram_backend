package server

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewServer(dbConnection *gorm.DB) *Server {
	return &Server{
		engine: gin.Default(),
		db:     dbConnection,
	}
}

var server *Server
func (server *Server) Run(addr string) error {
	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}

func (server *Server) Database() *gorm.DB {
	return server.db
}

// server instance
// var Serverr = socketio.NewServer(nil)
// func Namespaces() {

// 	// connecting client to the server
// 	Serverr.OnConnect("/", chat.Connect)
// 	Serverr.OnEvent("/", "createroom", chat.RoomCreate)
// 	Serverr.OnEvent("/", "join", chat.RoomJoin)
// 	Serverr.OnEvent("/", "message", chat.Messages)
// 	Serverr.OnEvent("/", "leave", chat.RoomLeave)
// 	Serverr.OnEvent("/", "voice", chat.VoiceCall)



// }
