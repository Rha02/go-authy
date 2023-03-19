package handlers

import "net/http"

// CreateToken() creates a new refresh token and a new access token
func (m *Repository) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Token"))
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
