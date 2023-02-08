package handler

import (
	"encoding/json"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (s *apiServer) NewComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {
		log.Println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check post is exist ?
	if booll := execute.CheckPostByid(s.DB, postID); !booll {
		w.Write([]byte("Not Found|404"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.Write([]byte("Bad Request|400"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var comment models.NewComment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.Write([]byte("Bad Request | 400"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert into database comment
	if res, booll := execute.CreateCommentSql(comment, postID, s.DB); !booll {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
