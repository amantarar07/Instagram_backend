package db

import (
	"main/server/model"

	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	
	
	
		err := db.AutoMigrate(&model.User{},&model.UserAuthSessions{},&model.Post{},&model.Liked_Posts{},&model.Viewed_Posts{},&model.Comments{})
		
		if err != nil {
			panic(err)
		}
	

}
