package handlers

import (
	"net/http"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
)

func (s *apiServer) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Log.LogError("Not correct method to delete")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !execute.DeleteToken(s.DB, tokenClient.Value) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "null",
		MaxAge:   -1,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}
