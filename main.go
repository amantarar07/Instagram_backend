package main

import (
	"log"
	"main/server"
	"main/server/db"
	auth "main/server/services/authentication"

	"os"

	"github.com/joho/godotenv"
)

// @title Instagram clone
// @version 1.0
// @description Social Media App
// @BasePath /localhost:3000

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := db.InitDB()
	db.Transfer(connection)


	//twilio
	auth.TwilioInit(os.Getenv("TWILIO_AUTH_TOKEN"))




	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
