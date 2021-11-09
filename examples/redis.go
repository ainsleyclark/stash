// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package examples

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lacuna-seo/stash"
	"log"
	"time"
)

func Redis() {
	// Create a provider, this could be Redis, Memcache
	// or In Memory (go-cache).
	provider := stash.NewRedis(redis.Options{
		Addr: "127.0.0.1:6379",
	}, time.Hour * 8)

	// Create a new cache store by passing a provider.
	cache, err := stash.Load(provider)
	if err != nil {
		log.Fatalln(err)
	}

	// Set a cache item with key and value, with the expiration
	// time of one hour and a tag,
	err = cache.Set(context.Background(), "key", []byte("stash"), stash.Options{
		Expiration: time.Hour * 1,
		Tags:       []string{"tag"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Obtains the cache item by key which automatically unmarshalls
	// the value by passing a reference to the same type that
	// has been stored.
	var buf []byte
	err = cache.Get(context.Background(), "key", &buf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buf)) // Returns stash
}
