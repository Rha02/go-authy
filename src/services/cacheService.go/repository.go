package cacheService

// CacheRepository is the interface for the cache repository
type CacheRepository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
