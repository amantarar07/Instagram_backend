package request


type Filepath struct{
	File_path string `json:"file_path"`
}

type Comment struct{

	PostID string `json:"post_id"`
	Comment string `json:"comment"`
}