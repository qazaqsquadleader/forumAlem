package execute

import (
	"database/sql"
	"forum-backend/internal/Log"
	"forum-backend/internal/models"
)

func CheckPasswordSql(User models.CheckUser, db *sql.DB) (models.User, bool) {
	var fullUser models.User
	query := `SELECT * FROM user WHERE username=$1 and password=$2`
	row := db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email); err != nil {
		Log.LogError(err.Error())
		return fullUser, false
	}
	// userId := fullUser.Username

	// fullUser.Token = "Token"
	// fullUser.ExpiresAt = time.Now().AddDate(0, 0, 14).String()
	return fullUser, true
}
