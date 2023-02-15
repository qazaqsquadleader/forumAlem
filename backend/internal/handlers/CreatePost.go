package handlers

// import (
// 	"encoding/json"
// 	"forum-backend/internal/database/execute"
// 	"log"
// 	"net/http"
// )

// func (s *apiServer) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		w.Write([]byte("Allowed method is GET"))
// 		return
// 	}
// 	tokenClient, err := r.Cookie("token")
// 	// log.Println(tokenClient)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	allPost, err := execute.GetAllpostSql(s.DB)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(allPost)
// 	log.Println(allPost)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// }

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		Log.LogError("Couldn't read the body of a request in SignInHandler or body is empty")
		w.WriteHeader(400)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {

		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var post models.NewPost
	err = json.Unmarshal(body, &post)
	if err != nil {

		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CreatePostSql(post, s.DB); !booll {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postLog := fmt.Sprintf("New Post Created by: %s", tokenClient.Value)

	Log.LogInfo(postLog)
	w.WriteHeader(http.StatusCreated)
}
