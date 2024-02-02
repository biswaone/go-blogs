package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/biswaone/go-blogs/internal/users"
	"github.com/jackc/pgx/v5"
)

func LoginHandler(db *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "This method is not allowed.", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var dto users.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, "Invalid JSON in request: "+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := dto.ValidateLoginRequest(db)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := err.Error()
			response := users.LoginResponse{Message: errMessage}
			json.NewEncoder(w).Encode(response)
			return
		}

		token, err := dto.Login(*user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := err.Error()
			response := users.LoginResponse{Message: errMessage}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		response := users.LoginResponse{Message: "User Loggedin Successfully", AccessToken: token}
		json.NewEncoder(w).Encode(response)

	})
}
