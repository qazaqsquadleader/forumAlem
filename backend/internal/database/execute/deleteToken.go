package execute

import (
	"database/sql"
	"fmt"

	"forum-backend/internal/Log"
)

func DeleteToken(db *sql.DB, clientToken string) bool {
	fmt.Println(clientToken)
	query := `DELETE FROM user_sessions WHERE token=$1`
	_, err := db.Exec(query, clientToken)

	// err := db.QueryRow(query, clientToken)
	fmt.Println(err)
	if err != nil {
		Log.LogError(err.Error())
		return false
	}
	return true
}
