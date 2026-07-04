package packagedesign

// Cache stores string values behind a small behavior-focused API.
type Cache struct {
	items map[string]string
}

// NewCache constructs a cache with hidden storage details.
func NewCache() *Cache {
	return &Cache{items: make(map[string]string)}
}

// Set records a value for key.
func (c *Cache) Set(key, value string) {
	c.items[key] = value
}

// Get returns the cached value and whether it exists.
func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.items[key]
	return value, ok
}
