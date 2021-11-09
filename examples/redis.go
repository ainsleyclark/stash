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
)

func Redis() {
	provider := stash.NewRedis(redis.Options{
		Addr: "127.0.0.1",
	})

	cache, err := stash.Load(provider)
	if err != nil {
		return
	}

	var buf []byte
	err = cache.Get(context.Background(), "key", buf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buf))
}
