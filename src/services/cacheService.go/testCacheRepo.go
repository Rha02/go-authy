package cacheService

type TestCacheRepo struct {
	store map[string]string
}

func NewTestCacheRepo() CacheRepository {
	return &TestCacheRepo{
		store: make(map[string]string),
	}
}

// Get() gets the value of the key from the cache
func (c *TestCacheRepo) Get(key string) (string, error) {
	return c.store[key], nil
}

// Set() sets the value of the key in the cache
func (c *TestCacheRepo) Set(key string, value string) error {
	c.store[key] = value
	return nil
}

// Delete() deletes the key from the cache
func (c *TestCacheRepo) Delete(key string) error {
	delete(c.store, key)
	return nil
}
