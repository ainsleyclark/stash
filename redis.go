// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"context"
	"errors"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisStore defines the data stored for the redisStore
// client.
type RedisStore struct {
	client  *redis.Client
	options redis.Options
	defaultExpiration time.Duration
}

// NewRedis creates a new redis store and returns a provider.
func NewRedis(options redis.Options, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{
		client:  redis.NewClient(&options),
		options: options,
		defaultExpiration: defaultExpiration,
	}
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (r *RedisStore) Validate() error {
	if r.options.Addr == "" {
		return errors.New("error: no redis address defined")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory driver name.
func (r *RedisStore) Driver() string {
	return RedisDriver
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (r *RedisStore) Store() store.StoreInterface {
	return cache.New(store.NewRedis(r.client, &store.Options{
		Expiration: r.defaultExpiration,
	}))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (r *RedisStore) Ping() error {
	return r.client.Ping(context.Background()).Err()
}
