// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stasher

import "github.com/eko/gocache/v2/store"

// StoreInterface used for mocking.
type StoreInterface interface {
	store.StoreInterface
}

// Provider used for mocking.
type Provider interface {
	Ping() error
	Validate() error
	Driver() string
	Store() store.StoreInterface
}
