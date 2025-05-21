package models

import "testing"

func TestRegisterValidate(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		username string
		password string
		expected string
	}{
		{
			name:     "Valid email, username and password",
			email:    "test@test.com",
			username: "test",
			password: "password",
			expected: "",
		},
		{
			name:     "Invalid email",
			email:    "test",
			username: "test",
			password: "password",
			expected: "Invalid email",
		},
		{
			name:     "Invalid password",
			email:    "test@test.com",
			username: "test",
			password: "",
			expected: "Invalid password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := RegisterRequest{
				Email:    tc.email,
				Name:     tc.username,
				Password: tc.password,
			}

			result := rr.Validate()
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}

}
