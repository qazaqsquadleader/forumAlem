package models

type NewUser struct {
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"password"`
}

type CheckUser struct {
	Username string `json:"Username"`
	Password string `json:"password"`
}

type User struct {
	UserId   int    `json:"UserId"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type NewPost struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"Content"`
}

type NewComment struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

type Like struct {
	CommentID     int  `json:"comment_id"`
	LikeOrDislike bool `json:"like"`
}
