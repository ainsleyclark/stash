// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"errors"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	pkg "github.com/go-redis/redis/v8"
)

// redis defines the data stored for the redis
// client.
type redis struct {
	client *pkg.Client
	config Config
}

// init adds the redis store to the providerMap
// on initialisation of the app.
func init() {
	providers.RegisterProvider(RedisStore, func(cfg Config) provider {
		return &redis{pkg.NewClient(&cfg.RedisOptions), cfg}
	})
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (r *redis) Validate() error {
	if r.config.RedisOptions.Addr == "" {
		return errors.New("no redis address defined in env")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory driver name.
func (r *redis) Driver() string {
	return RedisStore
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (r *redis) Store() store.StoreInterface {
	return cache.New(store.NewRedis(r.client, options))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (r *redis) Ping() error {
	return r.client.Ping(context.Background()).Err()
}
