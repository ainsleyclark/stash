// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/eko/gocache/v2/store"
)

// Provider defines the methods for a cache Provider.
type Provider interface {
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
