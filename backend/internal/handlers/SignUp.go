package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"forum-backend/internal/Log"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		Log.LogError(err.Error())
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var usr models.NewUser
	err = json.Unmarshal(body, &usr)
	if err != nil {
		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userCreated := execute.CreateUserSql(usr, s.DB); !userCreated {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Log.LogInfo("Successfully created a user")
	w.WriteHeader(http.StatusCreated)
}
