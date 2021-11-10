// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import "time"

func (t *StashTestSuite) TestMemory() {
	got := NewMemory(time.Second*1, time.Second*1)
	t.NotNil(got)
	m := memoryStore{}
	t.Nil(m.Validate())
	t.Equal(MemoryDriver, m.Driver())
	store := m.Store()
	t.NotNil(store)
	t.Nil(m.Ping())
}
