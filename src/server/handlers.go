package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AIPCB/auth-service/src/models"
)

// TODO: Prevent duplicate records
func (s *Server) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		errorMsg := req.Validate()
		if errorMsg != "" {
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		user, err := s.storage.CreateUser(r.Context(), req.Email, req.Username, req.Password)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		response := models.RegisterResponse{
			Message: fmt.Sprintf("Successfully registered user %s", user.Username),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		errorMsg := req.Validate()
		if errorMsg != "" {
			http.Error(w, errorMsg, http.StatusBadRequest)
			return
		}

		// todo: implement actual logic for login
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(req)
	}
}
