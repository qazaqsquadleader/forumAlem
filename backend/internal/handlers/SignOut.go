package handlers

import (
	"fmt"
	"net/http"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
)

func (s *apiServer) SignOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tokenClient, err := r.Cookie("token")
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tokenClient.Value == "" {
		Log.LogError("User's token has no value of type token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if exists := execute.DeleteToken(s.DB, tokenClient.Value); !exists {
		Log.LogError("User's token isn't present in the DB")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println("Using setCookie")
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "null",
			MaxAge:   -1,
			HttpOnly: false,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}

	Log.LogInfo(fmt.Sprintf("User with id %s ended session", tokenClient.Value))
	w.WriteHeader(http.StatusOK)
}
