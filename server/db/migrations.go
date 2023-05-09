package db

import (
	"main/server/model"

	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	
	
	
		err := db.AutoMigrate(&model.User{},&model.UserAuthSessions{},&model.Post{},&model.Comment{},&model.Like{},model.Participants{},model.Room{},model.Message{})
		
		if err != nil {
			panic(err)
		}
	

}
