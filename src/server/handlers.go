package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AIPCB/auth-service/src/models"
)

// TODO: Prevent duplicate records
func (s *Server) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		var req models.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		errorMsg := req.Validate()
		if errorMsg != "" {
			http.Error(w, errorMsg, http.StatusBadRequest)
			log.Printf("Validation error: %s", errorMsg)
			return
		}

		err = s.personService.CreatePerson(ctx, req)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Error creating person: %v", err)
			return
		}

		token, err := s.GenerateToken()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Error generating token: %v", err)
			return
		}

		response := models.RegisterResponse{
			Message:     fmt.Sprintf("Successfully registered user"),
			Success:     true,
			AccessToken: token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

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

		person, err := s.personService.GetPersonByEmail(ctx, req.Email)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Error fetching person by email: %v", err)
			return
		}

		token, err := s.GenerateToken()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Error generating token: %v", err)
			return
		}

		response := models.LoginResponse{
			Message:     fmt.Sprintf("Successfully logged in user %s", person.Name),
			AccessToken: token,
			Success:     true,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
