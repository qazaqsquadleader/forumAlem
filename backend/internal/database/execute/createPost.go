package execute

import (
	"database/sql"
	"time"

	"forum-backend/internal/Log"
	"forum-backend/internal/models"
)

func CreatePostSql(post models.NewPost, db *sql.DB) bool {
	stmt, err := db.Prepare("INSERT INTO posts(author, title,content,creationDate) values(?,?,?,?)")
	if err != nil {
		Log.LogError(err.Error())
		return false
	}
	if _, err := stmt.Exec(post.Author, post.Title, post.Content, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		Log.LogError(err.Error())
		return false
	}
	return true
}
