package routes

import (
	"net/http"

	"github.com/Rha02/go-authy/src/handlers"
	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create-token", handlers.Repo.CreateToken).Methods("POST")
	router.HandleFunc("/verify-token", handlers.Repo.VerifyToken).Methods("POST")
	router.HandleFunc("/refresh-token", handlers.Repo.RefreshToken).Methods("POST")
	router.HandleFunc("/revoke-token", handlers.Repo.RevokeToken).Methods("POST")

	return router
}
