package execute

import (
	"database/sql"
	"forum-backend/internal/models"
)

func GetByToken(db *sql.DB, clientToken string) (models.User, bool, error) {
	var User models.User
	var id int
	query := `SELECT userId FROM user_sessions WHERE token=$1`
	err := db.QueryRow(query, clientToken).Scan(&id)
	if err == sql.ErrNoRows {
		return User, false, err
	}
	if err != nil {
		return User, false, err
	}
	query1 := `SELECT * FROM user WHERE userId=$1`
	row := db.QueryRow(query1, id)
	if err := row.Scan(&User.UserId, &User.Username, &User.Password, &User.Email); err != nil {
		return models.User{}, false, err
	}

	return User, true, nil
}
