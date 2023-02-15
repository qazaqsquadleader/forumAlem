package handlers

import (
	"net/http"

	handlers "forum-backend/internal/handlers/functionality"
)

func (s *apiServer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// isAuthenticated, ok := r.Context().Value("isAuthenticated").(bool)
	w.WriteHeader(http.StatusOK)
	handlers.GetPosts(s.DB, w)
	// tokenClient, err := r.Cookie("token")
	// log.Println(tokenClient)
	// if err != nil {
	// 	log.Println(err.Error)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Bad Request"))
	// 	return
	// }
}
