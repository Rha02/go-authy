package main

import (
	"fmt"
	"net/http"

	"github.com/Rha02/go-authy/src/handlers"
	"github.com/Rha02/go-authy/src/routes"
	"github.com/Rha02/go-authy/src/services/cacheService.go"
)

const port = "8080"

func main() {
	// Initialize the cache repo
	cacheRepo := cacheService.NewRedisRepo()

	// Initialize the handlers repo
	repo := handlers.NewRepository(cacheRepo)
	handlers.NewHandlers(repo)

	// Create the router
	router := routes.Routes()

	// Start the server
	server := &http.Server{
		Addr:    fmt.Sprint(":", port),
		Handler: router,
	}

	fmt.Printf("Server started on port %s", port)
	server.ListenAndServe()
}
