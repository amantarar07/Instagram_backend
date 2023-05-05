package model

import "gorm.io/gorm"



type Post struct {
	gorm.Model
	
	PostID string `json:"post_id" gorm:"default:uuid_generate_v4();unique"`
	Title   string `json:"title"`
	Caption  string `json:"caption"`
	Path string `json:"path"`
	UserID  string `json:"user_id"`
	Likes int64 `json:"likes"`
	Comments int64 `json:"comments"`
	Views  int64 `json:"views"`
	
	
}

type Like struct{

	gorm.Model
	LikeID  string `json:"like_id" gorm:"default:uuid_generate_v4();unique"`
	PostID string `json:"post_id"`
	CommentID string `json:"comment_id"`
	User_Id string `json:"user_id"`

}

// type Commented_Posts struct{

// 	//gorm.Model

// 	PostID string `json:"post_id"`
// 	Who_commented string `json:"who_commented"`
// 	Comment []string `gorm:"type:text[]"`
	
// }

// type Viewed_Posts struct{

// 	gorm.Model

// 	PostID string `json:"post_id"`
// 	Who_viewed string `json:"who_viewed"`
// }

type Comment struct {

	CommentID string `json:"comment_id" gorm:"default:uuid_generate_v4();unique"` 
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	CommentText string `json:"comment_text"`
	LikesCount int64 `json:"likes_count"`

}