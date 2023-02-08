package execute

import (
	"database/sql"
)

func CheckPostByid(db *sql.DB, id int) bool {
	selectRecord := "SELECT title FROM user_sessions WHERE posts = ?"
	var title string
	err := db.QueryRow(selectRecord, id).Scan(&title)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		return false
	}
	return true
}
