package model 


type Followers struct{

	User_id string `json:"user_id"`
	FollowerOf  string `json:"follower"`

}

// type Following struct{

// 	User_id string `json:"user_id"`
// 	FollowingWhom string `json:"following_whom"`
// }