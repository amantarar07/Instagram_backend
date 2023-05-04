package model

import "gorm.io/gorm"



type Post struct {
	gorm.Model
	
	PostID string `json:"post_id" gorm:"default:uuid_generate_v4();unique"`
	Title   string `json:"title"`
	Path string `json:"path"`
	UserID  string `json:"user_id"`
	Likes int64 `json:"likes"`
	CommentCount int64 `json:"comment_count"`
	Comment []string `gorm:"type:text[]"`
	Views  int64 `json:"views"`
	
	
}

type Liked_Posts struct{

	gorm.Model

	PostID string `json:"post_id"`
	Who_liked string `json:"who_liked"`

}

// type Commented_Posts struct{

// 	//gorm.Model

// 	PostID string `json:"post_id"`
// 	Who_commented string `json:"who_commented"`
// 	Comment []string `gorm:"type:text[]"`
	
// }

type Viewed_Posts struct{

	gorm.Model

	PostID string `json:"post_id"`
	Who_viewed string `json:"who_viewed"`
}

type Comments struct {


	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Comment string 

}