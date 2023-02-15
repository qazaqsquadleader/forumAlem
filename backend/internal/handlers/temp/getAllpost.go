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

// func (s *apiServer) GetAllpost(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.Write([]byte("Allowed method is GET"))
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// 	tokenClient, err := r.Cookie("token")
// 	log.Println(tokenClient)
// 	if err != nil {
// 		_, fileName, lineNum, _ := runtime.Caller(0)
// 		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
// 		s.log.Output(errStr)
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
// 		_, fileName, lineNum, _ := runtime.Caller(0)
// 		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
// 		s.log.Output(errStr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(allPost)
// 	log.Println(allPost)
// 	if err != nil {
// 		_, fileName, lineNum, _ := runtime.Caller(0)
// 		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
// 		s.log.Output(errStr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	return
// }
