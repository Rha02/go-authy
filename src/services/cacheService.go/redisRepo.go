package cacheService

// RedisRepo is a Redis cache implementation of the CacheRepository interface
type RedisRepo struct{}

func NewRedisRepo() CacheRepository {
	return &RedisRepo{}
}

// Get() gets the value of the key from the cache
func (c *RedisRepo) Get(key string) (string, error) {
	return "", nil
}

// Set() sets the value of the key in the cache
func (c *RedisRepo) Set(key string, value string) error {
	return nil
}

// Delete() deletes the key from the cache
func (c *RedisRepo) Delete(key string) error {
	return nil
}
