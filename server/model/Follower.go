package model 


type Followers struct{

	Follower string `json:"follower"`
	FollowerOf  string `json:"followerof"`

}

// type Following struct{

// 	User_id string `json:"user_id"`
// 	FollowingWhom string `json:"following_whom"`
// }