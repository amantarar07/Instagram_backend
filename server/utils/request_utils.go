package utils

import (
	"encoding/json"
	"io/ioutil"
	"main/server/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RequestDecoding(context *gin.Context, data interface{}) {

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
}

func SetHeader(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
