package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
)

func GetPosts(DB *sql.DB, w http.ResponseWriter) {
	allPost, err := execute.GetAllpostSql(DB)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	err = json.NewEncoder(w).Encode(allPost)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
}
