// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"errors"
	pkg "github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
)

// memcache defines the data stored for the memcache
// client.
type memcache struct {
	client *pkg.Client
	config Config
}

// init adds the memcache store to the providerMap
// on initialisation of the app.
func init() {
	providers.RegisterProvider(MemcacheStore, func(cfg Config) provider {
		return &memcache{pkg.New(cfg.MemCachedHosts), cfg}
	})
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (m *memcache) Validate() error {
	if m.config.MemCachedHosts == "" {
		return errors.New("no memcache hosts defined in env")
	}
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory driver name.
func (m *memcache) Driver() string {
	return MemcacheStore
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (m *memcache) Store() store.StoreInterface {
	return cache.New(store.NewMemcache(m.client, options))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (m *memcache) Ping() error {
	return m.client.Ping()
}
