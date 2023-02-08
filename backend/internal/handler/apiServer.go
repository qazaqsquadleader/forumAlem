package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type apiServer struct {
	DB     *sql.DB
	Router *http.ServeMux
}

func NewApiServer(db *sql.DB) *apiServer {
	return &apiServer{
		DB:     db,
		Router: http.NewServeMux(),
	}
}

func CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		fmt.Println("Cors applied")
		next.ServeHTTP(w, r)
	})
}

func (s *apiServer) Start() error {
	log.Println("Starting the server at port localhost:8080")

	s.Router.Handle("/api/checkPassword", CorsHeaders(http.HandlerFunc(s.CheckPassword)))
	s.Router.Handle("/api/createUser", CorsHeaders(http.HandlerFunc(s.CreateUser)))
	s.Router.Handle("/api/createPost", CorsHeaders(http.HandlerFunc(s.CreatePost)))
	s.Router.HandleFunc("/api/createComment", s.NewComment)
	s.Router.HandleFunc("/api/newLike", s.Like)
	return http.ListenAndServe(":8080", s.Router)
}
