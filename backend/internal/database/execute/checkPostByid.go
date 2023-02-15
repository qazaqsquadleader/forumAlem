package execute

import (
	"database/sql"
	"forum-backend/internal/Log"
)

func CheckPostByid(db *sql.DB, id int) bool {
	selectRecord := "SELECT title FROM user_sessions WHERE posts = ?"
	var title string
	err := db.QueryRow(selectRecord, id).Scan(&title)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		Log.LogError(err.Error())
		return false
	}
	return true
}
