// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package examples

import (
	"context"
	"fmt"
	"github.com/lacuna-seo/stash"
	"log"
	"time"
)

// Memory example for Stash.
func Memory() {
	provider := stash.NewMemory(5*time.Minute, 10*time.Minute)

	cache, err := stash.Load(provider)
	if err != nil {
		log.Fatalln(err)
	}

	err = cache.Set(context.Background(), "key", []byte("stash"), stash.Options{
		Expiration: time.Hour * 1,
		Tags:       []string{"tag"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	var buf []byte
	err = cache.Get(context.Background(), "key", &buf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buf)) // Returns stash
}
