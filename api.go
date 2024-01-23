package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/biswaone/go-blogs/user"
	"github.com/jackc/pgx/v5"
)

func RegisterUserHandler(db *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		var dto user.RegisterUser
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, "Invalid JSON in request: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = dto.ValidateNewUser()
		if err != nil {
			http.Error(w, "Error: validation Error, Mssg: "+err.Error(), http.StatusBadRequest)
			return
		}
		newUser, err := dto.CreateUser(db)
		if err != nil {
			http.Error(w, "Cannot create user "+err.Error(), http.StatusBadRequest)
			return
		}
		response := user.Response{Message: "User Successfully Registered", User: *newUser}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
