// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

func (t *StashTestSuite) TestMemory() {
	got := NewMemory(DefaultExpiry, DefaultCleanup)
	t.NotNil(got)
	t.NotNil(got.client)
	m := MemoryStore{}
	t.Nil(m.Validate())
	t.Equal(MemoryDriver, m.Driver())
	store := m.Store()
	t.NotNil(store)
	t.Nil(m.Ping())
}
