package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Status  string      `json:"status"`
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(context *gin.Context, statusCode int, data interface{}) {
	context.JSON(statusCode, data)
}

func ShowResponse(status string, statusCode int64, message string, data interface{}, context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeader(int(statusCode))
	response := Success{
		Status:  status,
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	Response(context, int(statusCode), response)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, Error{Code: statusCode, Message: message})
}


type SocketResp struct{

	Message string `json:"message"`
	Data  interface{} `json:"data"`
}

func SocketResponse(data interface{}, message string, s socketio.Conn) {
	socketResponse := SocketResp{
		Message: message,
		Data:    data,
	}
	
	s.Emit("ack", socketResponse, func() {
		fmt.Println("acknowledgement sent to client")
	})
}
