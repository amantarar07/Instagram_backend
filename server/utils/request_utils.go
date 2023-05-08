package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/server/model"
	"main/server/response"

	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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


func IsEmail(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
}

func GenerateToken(claims model.Claims)string{

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


key:="secret"
ss, err := token.SignedString([]byte(key))

if err!=nil{

	fmt.Println("token not signed successfully")
}

return ss


}
func DecodeToken(tokenString string) (*model.Claims,error){

	fmt.Println("decode token called")
	claims:=&model.Claims{}
	//parse the token to get the claims 
	token, err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err!=nil{
		fmt.Println("err",err)
	}
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {

		return claims, nil
	} else {
		fmt.Println(ok)
		return nil,err
	}

}

