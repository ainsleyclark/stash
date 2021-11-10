// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stasher

import "github.com/eko/gocache/v2/store"

type StoreInterface interface {
	store.StoreInterface
}

type Provider interface {
	Ping() error
	Validate() error
	Driver() string
	Store() store.StoreInterface
}
