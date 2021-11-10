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

// redisStore defines the data stored for the redisStore
// client.
type redisStore struct {
	client            *redis.Client
	options           redis.Options
	defaultExpiration time.Duration
}

// NewRedis creates a new redis store and returns a provider.
func NewRedis(options redis.Options, defaultExpiration time.Duration) Provider {
	return &redisStore{
		client:            redis.NewClient(&options),
		options:           options,
		defaultExpiration: defaultExpiration,
	}
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (r *redisStore) Validate() error {
	if r.options.Addr == "" {
		return errors.New("error: no redis address defined")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory Driver name.
func (r *redisStore) Driver() string {
	return RedisDriver
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (r *redisStore) Store() store.StoreInterface {
	return cache.New(store.NewRedis(r.client, &store.Options{
		Expiration: r.defaultExpiration,
	}))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (r *redisStore) Ping() error {
	return r.client.Ping(context.Background()).Err()
}
