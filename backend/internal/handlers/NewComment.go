package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) NewComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		// _, ok := runtime.Caller(1)
		// if !ok {
		// 	log.Println("failed to get the runtime caller for the Logger")
		// }
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {
		// _, ok := runtime.Caller(1)
		// if !ok {
		// 	log.Println("failed to get the runtime caller for the Logger")
		// }
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check post is exist ?
	if booll := execute.CheckPostByid(s.DB, postID); !booll {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var comment models.NewComment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert into database comment
	if booll := execute.CreateCommentSql(comment, postID, s.DB); !booll {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
