// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/eko/gocache/v2/store"
	"time"
)

// Options represents the cache store available options
// when using Set().
type Options struct {
	// Expiration allows to specify a global expiration
	// time hen setting a value.
	Expiration time.Duration
	// Tags allows specifying associated tags to the
	// current value.
	Tags []string
}

// InvalidateOptions represents the options for invalidating
// the cache.
type InvalidateOptions struct {
	// Tags allows to specify associated tags to the
	// current value.
	Tags []string
}

// toStore converts Options to the store Options.
func (o *Options) toStore() *store.Options {
	return &store.Options{
		Expiration: o.Expiration,
		Tags:       o.Tags,
	}
}

// toStore converts InvalidateOptions to the store
// InvalidateOptions.
func (i *InvalidateOptions) toStore() store.InvalidateOptions {
	return store.InvalidateOptions{Tags: i.Tags}
}
