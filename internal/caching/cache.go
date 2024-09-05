package caching

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration)
}

type InMemoryCache struct {
	cache *cache.Cache
}

func NewInMemoryCache(defaultExpiration, cleanupInterval time.Duration) *InMemoryCache {
	return &InMemoryCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (c *InMemoryCache) Get(ctx context.Context, key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *InMemoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	c.cache.Set(key, value, expiration)
}
