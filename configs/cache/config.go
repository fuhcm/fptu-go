package cache

import (
	"time"

	"fptugo/pkg/caching"
)

// Create a cache
var c = caching.New(30*time.Minute, 60*time.Minute)

// GetCache ...
func GetCache() *caching.Cache {
	return c
}

// GetDefaultExpiration ...
func GetDefaultExpiration() time.Duration {
	return caching.DefaultExpiration
}
