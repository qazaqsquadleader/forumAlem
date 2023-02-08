package execute

import (
	"database/sql"
	"fmt"
	"forum-backend/internal/models"
)

func CreateCommentSql(newComment models.NewComment, id int, db *sql.DB) (string, bool) {
	stmt, err := db.Prepare("INSERT INTO comments(postId, author,content) values(?,?,?)")
	if err != nil {
		return "SQL INJECTION", false
	}
	if _, err := stmt.Exec(id, newComment.Author, newComment.Body); err != nil {
		fmt.Println(err.Error())
		return "Error with creation of new comment", false
	}
	return "", true
}
