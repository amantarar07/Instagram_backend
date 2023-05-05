package request


type Filepath struct{
	File_path string `json:"file_path"`
}

type Comment struct{

	PostID string `json:"post_id"`
	Comment string `json:"comment"`
}

type Like struct{

	PostID string `json:"post_id"`
	CommentID string `json:"comment_id"`
}

type Caption struct {

	CaptionText string `json:"caption"`
}

type Bio struct {

	Bio string `json:"bio"`
}