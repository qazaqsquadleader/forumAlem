package execute

import (
	"database/sql"

	"forum-backend/internal/Log"

	"forum-backend/internal/models"
)

func CreateUserSql(User models.NewUser, db *sql.DB) bool {
	stmt, err := db.Prepare("INSERT INTO User(username, password,email) values(?,?,?)")
	if err != nil {

		Log.LogError(err.Error())
		return false
	}
	if _, err := stmt.Exec(User.Username, User.Password, User.Email); err != nil {

		Log.LogError(err.Error())
		return false
	}
	return true
}
