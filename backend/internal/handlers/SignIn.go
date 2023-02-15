package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"

	"github.com/google/uuid"
)

func (s *apiServer) SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			Log.LogError("Couldn't read the body of a request in SignInHandler or body is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var usr models.CheckUser
		err = json.Unmarshal(body, &usr)
		if err != nil {
			Log.LogError(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if userModel, booll := execute.CheckPasswordSql(usr, s.DB); booll {
			sessionExists, sessionErr := UserSessionsExist(s.DB, userModel.UserId)
			if sessionErr != nil {
				Log.LogError(sessionErr.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if sessionExists {
				_, err := DeleteUserSessions(s.DB, userModel.UserId)
				if err != nil {
					Log.LogError(err.Error())
					return
				}
			}
			// 	if sessionNotExists {
			// 		Log.LogInfo("Creating a new session...")
			tokenStat, err := s.DB.Prepare(`INSERT INTO user_sessions (token, expiresAt, userId) VALUES (?, ?, ?);`)
			if err != nil {
				Log.LogError(err.Error())
				return
			}
			expiresAt := time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05")
			token := generateToken()
			_, err = tokenStat.Exec(token, expiresAt, userModel.UserId)
			if err != nil {
				Log.LogError(err.Error())
				return
			}
			cookie := &http.Cookie{
				Name:     "token",
				Value:    token,
				Expires:  time.Now().Add(time.Hour),
				HttpOnly: false,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
			loginLog := fmt.Sprintf("User with token: %s, just logged in...", token)
			Log.LogInfo(loginLog)
			err = json.NewEncoder(w).Encode(userModel)
			if err != nil {
				Log.LogError(err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}
		}
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func generateToken() string {
	token := uuid.New().String()
	tokenStr := fmt.Sprintf("Token created: %s", token)
	Log.LogInfo(tokenStr)
	return token
}

//	func sessionNotExists(db *sql.DB, userID int) (bool, error) {
//		// check if the token already exists in the sessions table
//		selectRecord := "SELECT token FROM user_sessions WHERE userId = ?"
//		var token string
//		err := db.QueryRow(selectRecord, userID).Scan(&token)
//		if err == sql.ErrNoRows {
//			// Handle case where no token exists for provided userId
//			Log.LogInfo("Session doesn't exist")
//			return true, nil
//		} else if err != nil {
//			// Handle other errors
//			return false, err
//		}
//		return false, nil
//	}
func UserSessionsExist(db *sql.DB, userID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM user_sessions WHERE userId=$1)`
	err := db.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		Log.LogError(err.Error())
		return false, err
	}
	return exists, nil
}

func DeleteUserSessions(db *sql.DB, userID int) (bool, error) {
	query := `DELETE FROM user_sessions WHERE userId=$1`
	_, err := db.Exec(query, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}
