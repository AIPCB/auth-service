package server

import (
	"encoding/json"
	"net/http"

	"github.com/AIPCB/auth-service/src/models"
)

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

		// todo: implement actual logic for registration
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(req)
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
