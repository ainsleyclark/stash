// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"context"
	"errors"
	"fmt"
	"github.com/lacuna-seo/stash/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	CacheKey = "key"
)

// StashTestSuite defines the helper used for cache
// testing.
type StashTestSuite struct {
	suite.Suite
}

// TestStash asserts testing has begun.
func TestStash(t *testing.T) {
	suite.Run(t, new(StashTestSuite))
}

// Setup assigns a mock Store to c.
func (t *StashTestSuite) Setup(mf func(m *mocks.StoreInterface)) *Cache {
	m := &mocks.StoreInterface{}
	if mf != nil {
		mf(m)
	}
	return &Cache{
		store: m,
	}
}

func (t *StashTestSuite) TestLoad() {
	tt := map[string]struct {
		mock func(m *mocks.Provider)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(nil)
				m.On("Ping").Return(nil)
				m.On("Driver").Return(MemoryDriver)
				m.On("Store").Return(nil)
			},
			MemoryDriver,
		},
		"Validate Error": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(errors.New("validate error"))
			},
			"validate error",
		},
		"Ping Error": {
			func(m *mocks.Provider) {
				m.On("Validate").Return(nil)
				m.On("Ping").Return(errors.New("ping error"))
				m.On("Driver").Return(MemoryDriver)
			},
			"ping error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Provider{}
			if test.mock != nil {
				test.mock(m)
			}
			c, err := Load(m)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			if c == nil {
				t.Fail("nil Driver")
				return
			}
			t.Equal(test.want, c.Driver)
		})
	}
}

func (t *StashTestSuite) TestLoad_NilProvider() {
	_, got := Load(nil)
	t.Contains(got.Error(), "provider cannot be nil")
}

func (t *StashTestSuite) TestStash_Get() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		run  func(cache *Cache) (interface{}, error)
		want interface{}
	}{
		"String": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).
					Return("\"item\"", nil)
			},
			func(c *Cache) (interface{}, error) {
				var tmp string
				err := c.Get(context.Background(), "key", &tmp)
				return tmp, err
			},
			"item",
		},
		"Int": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).
					Return("1", nil)
			},
			func(c *Cache) (interface{}, error) {
				var tmp int
				err := c.Get(context.Background(), "key", &tmp)
				return tmp, err
			},
			1,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).
					Return(nil, fmt.Errorf("get error"))
			},
			func(c *Cache) (interface{}, error) {
				var tmp string
				err := c.Get(context.Background(), "key", &tmp)
				return tmp, err
			},
			"get error",
		},
		"Byte Slice": {
			func(m *mocks.StoreInterface) {
				m.On("Get", mock.Anything, mock.Anything).
					Return([]byte("\"test\""), nil)
			},
			func(c *Cache) (interface{}, error) {
				var tmp string
				err := c.Get(context.Background(), "key", &tmp)
				return tmp, err
			},
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got, err := test.run(c)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StashTestSuite) TestStash_Set() {
	tt := map[string]struct {
		mock  func(m *mocks.StoreInterface)
		value interface{}
		want  interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			"key",
			nil,
		},
		"Marshal Error": {
			nil,
			make(chan bool),
			"json: unsupported type",
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(fmt.Errorf("set error"))
			},
			"key",
			"set error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got := c.Set(context.Background(), "key", test.value, Options{})
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StashTestSuite) TestStash_Delete() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Delete", mock.Anything, mock.Anything).
					Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Delete", mock.Anything, mock.Anything).
					Return(fmt.Errorf("delete error"))
			},
			"delete error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got := c.Delete(context.Background(), "key")
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StashTestSuite) TestStash_Invalidate() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Invalidate", mock.Anything, mock.Anything).
					Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Invalidate", mock.Anything, mock.Anything).
					Return(fmt.Errorf("invalidate error"))
			},
			"invalidate error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got := c.Invalidate(context.Background(), InvalidateOptions{})
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StashTestSuite) TestStash_Clear() {
	tt := map[string]struct {
		mock func(m *mocks.StoreInterface)
		want interface{}
	}{
		"Success": {
			func(m *mocks.StoreInterface) {
				m.On("Clear", mock.Anything).
					Return(nil)
			},
			nil,
		},
		"Error": {
			func(m *mocks.StoreInterface) {
				m.On("Clear", mock.Anything).
					Return(fmt.Errorf("clear error"))
			},
			"clear error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.mock)
			got := c.Clear(context.Background())
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
