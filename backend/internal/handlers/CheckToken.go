package handlers

import (
	"encoding/json"
	"net/http"

	"forum-backend/internal/Log"

	"forum-backend/internal/database/execute"
)

func (s *apiServer) CheckToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tokenClient, err := r.Cookie("token")
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if tokenClient.Value == "" {
		Log.LogInfo("No valid token with name 'token'")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Log.LogInfo(tokenClient.Value)

	User, booll, err := execute.GetByToken(s.DB, tokenClient.Value)
	if !booll {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
