package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
)

func (s *apiServer) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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
	res, booll, err := execute.GetByToken(s.DB, tokenClient.Value)
	if !booll || err != nil {
		Log.LogInfo("No session matches user token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idPost, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if booll := execute.DeletePost(s.DB, res.Username, idPost); !booll {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Log.LogInfo(fmt.Sprintf("Post with id %v deleted by %s", idPost, tokenClient.Value))
	w.WriteHeader(http.StatusOK)
}
