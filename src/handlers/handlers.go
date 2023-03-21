package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rha02/go-authy/src/models"
)

// CreateToken() creates a new access token
func (m *Repository) CreateToken(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the request body into a user object
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	// Create a response body
	type responseBody struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}

	// Marshal the response body into a json object
	res, _ := json.Marshal(&responseBody{
		RefreshToken: fmt.Sprintf("refresh_token_%s", user.GetId()),
		AccessToken:  fmt.Sprintf("access_token_%s", user.GetId()),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// VerifyToken() verifies the integrity of an access token
func (m *Repository) VerifyToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Verify Token"))
}

// RefreshToken() refreshes the access token
func (m *Repository) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Refresh Token"))
}

// RevokeToken() invalidates the refresh token and the access token
func (m *Repository) RevokeToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Revoke Token"))
}
