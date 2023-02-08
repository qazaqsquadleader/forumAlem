package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"

	"github.com/google/uuid"
)

func (s *apiServer) CheckPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Wrong Method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		log.Println("Empty body")
		w.Write([]byte("StatusBadRequest"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var usr models.CheckUser
	err = json.Unmarshal(body, &usr)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("StatusBadRequest"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userModel, booll := execute.CheckPasswordSql(usr, s.DB); booll {
		// jData, err := json.Marshal(userModel)
		// if err != nil {
		// 	log.Println("Couldn't Marshal userModel in checkPassword")
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		sessionNotExists, sessionErr := sessionNotExists(s.DB, userModel.UserId)
		if sessionErr != nil {
			log.Println(sessionErr.Error())
			w.Write([]byte("StatusBadRequest"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if sessionNotExists {
			tokenStat, err := s.DB.Prepare(`INSERT INTO user_sessions (token, expiresAt, userId) VALUES (?, ?, ?);`)
			// (SELECT userId FROM user WHERE username = ?)
			if err != nil {
				log.Println(err)
				return
			}
			expiresAt := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
			token := generateToken()

			_, err = tokenStat.Exec(token, expiresAt, userModel.UserId)
			if err != nil {
				log.Println(err)
				return
			}
			cookie := &http.Cookie{
				Name:     "token",
				Value:    token,
				HttpOnly: false,
			}
			http.SetCookie(w, cookie)
		}
		// w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(userModel)
		if err != nil {
			log.Println("Error Parsing JSON after confirmation of userCheck")
			w.Write([]byte("StatusBadRequest"))
			w.WriteHeader(http.StatusBadRequest)
		}
		// w.Write(jData)
		// w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println("User not found")
	return
}

func generateToken() string {
	token := uuid.New().String()
	fmt.Println("Generated token:", token)
	return token
}

func sessionNotExists(db *sql.DB, userID int) (bool, error) {
	// check if the token already exists in the sessions table
	selectRecord := "SELECT token FROM user_sessions WHERE userId = ?"
	var token string
	err := db.QueryRow(selectRecord, userID).Scan(&token)
	if err == sql.ErrNoRows {
		// Handle case where no token exists for provided userId
		fmt.Println("Not in sessions")
		return true, nil
	} else if err != nil {
		// Handle other errors
		return false, err
	}
	return false, nil
}
