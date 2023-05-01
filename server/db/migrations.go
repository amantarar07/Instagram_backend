package db

import (
	"main/server/model"

	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	
	
	
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			panic(err)
		}
	

}
