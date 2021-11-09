// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

func (t *StashTestSuite) TestMemcache() {
	servers := []string{"10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212"}

	got := NewMemcache(servers, time.Second*1)
	t.NotNil(got)
	t.Equal(servers, got.servers)
	t.Equal(time.Second*1, got.defaultExpiration)

	t.UtilTestProviderSuccess(&MemcacheStore{
		client:  memcache.New(""),
		servers: servers,
	}, MemcacheDriver)

	t.UtilTestProviderError(&MemcacheStore{
		client: memcache.New(""),
	})
}
