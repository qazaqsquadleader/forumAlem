package execute

import (
	"database/sql"
	"forum-backend/internal/Log"
	"forum-backend/internal/models"
)

func CreateCommentSql(newComment models.NewComment, id int, db *sql.DB) bool {
	stmt, err := db.Prepare("INSERT INTO comments(postId, author,content) values(?,?,?)")
	if err != nil {
		Log.LogError(err.Error())
		return false
	}
	if _, err := stmt.Exec(id, newComment.Author, newComment.Body); err != nil {
		Log.LogError(err.Error())
		return false
	}
	return true
}
