// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/eko/gocache/v2/store"
)

// provider defines the methods for a cache provider.
type provider interface {
	// Ping the store.
	Ping() error
	// Validate checks the environment for errors.
	Validate() error
	// Driver returns the store's name.
	Driver() string
	// Store returns the interface for use within
	// the cache.
	Store() store.StoreInterface
}

// providerMap defines the map of providerAdder functions
// defined by their name.
type providerMap map[string]providerAdder

// providerAdder is used to obtain a cache provider by
// injecting the configuration and returning a new
// provider type.
type providerAdder func(cfg Config) provider

var (
	// providers is the in memory collection of cache
	// providers.
	providers = providerMap{}
	// options are the default cache store options.
	options = &store.Options{
		Expiration: DefaultExpiry,
	}
)

// RegisterProvider adds a provider to the provider map.
// FullPath the provider already exists the function will
// panic.
func (p providerMap) RegisterProvider(name string, fn providerAdder) {
	if p.Exists(name) {
		panic("Error, duplicate cache provider: " + name)
		return
	}
	p[name] = fn
}

// Exists checks to see if a provider already exists
// in the map by name.
func (p providerMap) Exists(name string) bool {
	_, exists := p[name]
	return exists
}
