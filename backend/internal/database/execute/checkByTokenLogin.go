package execute

import (
	"database/sql"
	"forum-backend/internal/Log"
)

func CheckByTokenLogin(db *sql.DB, clientToken string) bool {
	var id int
	query := `SELECT userId FROM user_sessions WHERE token=$1`
	err := db.QueryRow(query, clientToken).Scan(&id)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		Log.LogError(err.Error())
		return false
	}

	return true
}
