package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/biswaone/go-blogs/internal/api/user"
	"github.com/jackc/pgx/v5"
)

func RegisterUserHandler(db *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var dto user.RegisterUser
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, "Invalid JSON in request: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = dto.ValidateNewUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := err.Error()
			response := user.Response{Message: "Invalid User Request", Exception: &errMessage}
			json.NewEncoder(w).Encode(response)
			return
		}
		response, err := dto.CreateUser(db)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := err.Error()
			response := user.Response{Message: "Cannot Create User", Exception: &errMessage}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
