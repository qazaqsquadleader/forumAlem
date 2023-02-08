package execute

import (
	"database/sql"
	"fmt"
	"forum-backend/internal/models"
	"time"
)

func CreatePostSql(post models.NewPost, db *sql.DB) (string, bool) {
	stmt, err := db.Prepare("INSERT INTO posts(author, title,content,creationDate) values(?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return "SQL INJECTION", false
	}

	if asd, err := stmt.Exec(post.Author, post.Title, post.Content, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		fmt.Println(asd)
		fmt.Println(err.Error())
		return "Error with creation of new user", false
	}
	return "", true
}
