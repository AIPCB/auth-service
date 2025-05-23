package models

import "testing"

func TestLoginValidate(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		expected string
	}{
		{
			name:     "Valid email and password",
			email:    "test@test.com",
			password: "password",
			expected: "",
		},
		{
			name:     "Invalid email",
			email:    "test",
			password: "password",
			expected: "Invalid email",
		},
		{
			name:     "Invalid password",
			email:    "test@test.com",
			password: "",
			expected: "Invalid password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := LoginRequest{
				Email:    tc.email,
				Password: tc.password,
			}

			result := rr.Validate()
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}

}
