// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

func (t *StashTestSuite) UtilTestProviderSuccess(p Provider, name string) {
	// Test Driver() method
	t.Equal(name, p.Driver())

	// Test Validate() method
	err := p.Validate()
	t.NoErrorf(err, "expecting Provider to pass validation")

	// Test Store() method
	store := p.Store()
	t.NotNil(store)
}

func (t *StashTestSuite) UtilTestProviderError(p Provider) {
	// Test Validate() method
	err := p.Validate()
	t.Errorf(err, "expecting Provider to fail validation")

	// Test Ping() method
	pingErr := p.Ping()
	t.Errorf(pingErr, "expecting Provider have ping error")
}
