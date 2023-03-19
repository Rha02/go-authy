package handlers

import "github.com/Rha02/go-authy/src/services/cacheService.go"

// Repo is a global variable that holds all the requests handlers
var Repo *Repository

// Repository is a repository for handlers
type Repository struct {
	CacheRepo cacheService.CacheRepository
}

// NewRepository() creates a
func NewRepository(cacheRepo cacheService.CacheRepository) *Repository {
	return &Repository{
		CacheRepo: cacheRepo,
	}
}

// NewHandlers() creates a new handlers
func NewHandlers(r *Repository) {
	Repo = r
}
