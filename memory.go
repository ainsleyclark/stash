// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

// memoryStore defines the data stored for the go-cache
// client.
type memoryStore struct {
	client *gocache.Cache
}

// NewMemory creates a new go-cache store and returns a provider.
func NewMemory(defaultExpiration, cleanupInterval time.Duration) Provider {
	return &memoryStore{
		client: gocache.New(defaultExpiration, cleanupInterval),
	}
}

// Validate satisfies the Provider interface by checking
// for environment variables.
func (m *memoryStore) Validate() error {
	return nil
}

// Driver satisfies the Provider interface by returning
// the memory Driver name.
func (m *memoryStore) Driver() string {
	return MemoryDriver
}

// Store satisfies the Provider interface by creating a
// new store.StoreInterface.
func (m *memoryStore) Store() store.StoreInterface {
	return cache.New(store.NewGoCache(m.client, nil))
}

// Ping satisfies the Provider interface by pinging the
// store.
func (m *memoryStore) Ping() error {
	return nil
}
