package models

import validator "github.com/exception-raised/validation-module/src"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr *LoginRequest) Validate() string {
	if !validator.IsValidEmail(rr.Email) {
		return "Invalid email"
	}

	if !validator.IsValidPassword(rr.Password, 8, 32, "") {
		return "Invalid password"
	}

	return ""
}
