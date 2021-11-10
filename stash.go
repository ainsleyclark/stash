// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eko/gocache/v2/store"
	"github.com/spf13/cast"
	"sync"
)

// Store defines methods for interacting with the
// caching system.
type Store interface {
	// Get retrieves a specific item from the cache by key.
	// Returns an error if the item could not be found
	// or unmarshalled.
	Get(ctx context.Context, key, v interface{}) error

	// Set set's a singular item in memory by key, value
	// and options (tags and expiration time).
	// Returns an error if the item could not be set.
	Set(ctx context.Context, key interface{}, value interface{}, options Options) error

	// Delete removes a singular item from the cache by
	// a specific key.
	// Returns an error if the item could not be deleted.
	Delete(ctx context.Context, key interface{}) error

	// Invalidate removes items from the cache via the
	// InvalidateOptions passed.
	// Returns an error if the cache could not be invalidated.
	Invalidate(ctx context.Context, options InvalidateOptions) error

	// Clear removes all items from the cache.
	// Returns an error.
	Clear(ctx context.Context) error
}

// Cache defines the methods for interacting with the
// cache layer.
type Cache struct {
	// store is the package store interface used for interacting
	// with the cache store.
	store store.StoreInterface
	// Driver is the current store being used, it can be
	// MemoryDriver, RedisDriver or MemcachedDriver.
	Driver string
}

const (
	// MemoryDriver is the Redis Driver, depicted
	// in the environment.
	MemoryDriver = "memory"
	// RedisDriver is the Redis Driver, depicted
	// in the environment.
	RedisDriver = "redis"
	// MemcacheDriver is the Memcached Driver, depicted
	// in the environment.
	MemcacheDriver = "memcache"
	// RememberForever is an alias for setting the
	// cache item to never be removed.
	RememberForever = -1
)

var (
	// Prevents data races when setting & getting cache
	// items.
	mtx = sync.Mutex{}
)

// Load initialises the cache store by the environment.
// It will load a Driver into memory ready for setting
// getting setting and deleting. Drivers supported are Memory
// Redis and MemCached.
// Returns ErrInvalidDriver if the Driver passed does not exist.
func Load(prov Provider) (*Cache, error) {
	if prov == nil {
		return nil, errors.New("provider cannot be nil")
	}

	err := prov.Validate()
	if err != nil {
		return nil, err
	}

	err = prov.Ping()
	if err != nil {
		return nil, err
	}

	return &Cache{
		store:  prov.Store(),
		Driver: prov.Driver(),
	}, nil
}

// Get retrieves a specific item from the cache by key.
// Returns an error if the item could not be found
// or unmarshalled.
func (c *Cache) Get(ctx context.Context, key, v interface{}) error {
	mtx.Lock()
	defer mtx.Unlock()

	result, err := c.store.Get(ctx, key)
	switch r := result.(type) {
	case []byte:
		err = json.Unmarshal(r, v)
	case string:
		err = json.Unmarshal([]byte(r), v)
	}

	if err != nil {
		return err
	}

	return nil
}

// Set set's a singular item in memory by key, value
// and options (tags and expiration time).
// Returns an error if the item could not be set.
func (c *Cache) Set(ctx context.Context, key interface{}, value interface{}, options Options) error {
	mtx.Lock()
	defer mtx.Unlock()
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.store.Set(ctx, key, marshal, options.toStore())
}

// Delete removes a singular item from the cache by
// a specific key.
// Returns an error if the item could not be deleted.
func (c *Cache) Delete(ctx context.Context, key interface{}) error {
	mtx.Lock()
	defer mtx.Unlock()
	return c.store.Delete(ctx, cast.ToString(key))
}

// Invalidate removes items from the cache via the
// InvalidateOptions passed.
// Returns an error if the cache could not be invalidated.
func (c *Cache) Invalidate(ctx context.Context, options InvalidateOptions) error {
	mtx.Lock()
	defer mtx.Unlock()
	return c.store.Invalidate(ctx, options.toStore())
}

// Clear removes all items from the cache.
// Returns an error.
func (c *Cache) Clear(ctx context.Context) error {
	mtx.Lock()
	defer mtx.Unlock()
	return c.store.Clear(ctx)
}
