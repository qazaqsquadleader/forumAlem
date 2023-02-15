package execute

import (
	"database/sql"

	"forum-backend/internal/Log"
)

func DeletePost(db *sql.DB, Author string, idPost int) bool {
	stmt, err := db.Prepare("DELETE FROM posts WHERE postId = $1 and author =$2")
	if err != nil {
		Log.LogError(err.Error())
		return false
	}
	if _, err := stmt.Exec(stmt, idPost, Author); err != nil {
		Log.LogError(err.Error())
		return false
	}
	return true
}
