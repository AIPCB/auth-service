package models

import (
	validator "github.com/exception-raised/validation-module/src"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: validate username
func (rr *RegisterRequest) Validate() string {
	if !validator.IsValidEmail(rr.Email) {
		return "Invalid email"
	}

	if !validator.IsValidPassword(rr.Password, 8, 32, "") {
		return "Invalid password"
	}

	return ""
}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
}
