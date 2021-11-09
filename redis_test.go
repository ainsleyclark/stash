// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/go-redis/redis/v8"
	"time"
)

func (t *StashTestSuite) TestRedis() {
	store := &redis.Options{Addr: "127.0.0.1", Password: ""}

	got := NewRedis(redis.Options{}, time.Second*1)
	t.NotNil(got)
	t.NotNil(got.options)
	t.NotNil(got.client)
	t.Equal(got.defaultExpiration, time.Second*1)

	t.UtilTestProviderSuccess(&RedisStore{
		client:  redis.NewClient(store),
		options: redis.Options{Addr: "127.0.0.1", Password: ""},
	}, RedisDriver)

	t.UtilTestProviderError(&RedisStore{
		client:  redis.NewClient(store),
		options: redis.Options{},
	})
}
