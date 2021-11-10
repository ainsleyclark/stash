# Stash
A cross driver cache store (stash) for Redis, MemCached & In-Memory storage. Stash is a wrapper for [GoCache](https://github.com/eko/gocache)
with automatic marshaling and unmarshalling of cache items.

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Test](https://github.com/lacuna-seo/stash/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lacuna-seo/stash/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lacuna-seo/stash/branch/master/graph/badge.svg?token=K27L8LS7DA)](https://codecov.io/gh/lacuna-seo/stash)
[![GoReportCard](https://goreportcard.com/badge/github.com/lacuna-seo/stash)](https://goreportcard.com/report/github.com/lacuna-seo/stash)
[![GoDoc](https://godoc.org/github.com/lacuna-seo/stash?status.png)](https://godoc.org/github.com/lacuna-seo/stash)

## Built-in stores

* [Memory (go-cache)](https://github.com/patrickmn/go-cache) (patrickmn/go-cache)
* [Memcache](https://github.com/bradfitz/gomemcache) (bradfitz/memcache)
* [Redis](https://github.com/go-redis/redis/v8) (go-redis/redis)

## Install

`go get -u github.com/lacuna-seo/stash`

## Provider

A provider can 

## Store

```go
// Store defines methods for interacting with the
// caching system.
type Store interface {
	// Get retrieves a specific item from the cache by key.
	// Returns an error if the item could not be found
	// or unmarshalled.
	Get(ctx context.Context, key, v interface{}) error

	// Set set's a singular item in memory by key, value
	// and options (tags and expiration time).
	// Returns an error if the item could not be set.
	Set(ctx context.Context, key interface{}, value interface{}, options Options) error

	// Delete removes a singular item from the cache by
	// a specific key.
	// Returns an error if the item could not be deleted.
	Delete(ctx context.Context, key interface{}) error

	// Invalidate removes items from the cache via the
	// InvalidateOptions passed.
	// Returns an error if the cache could not be invalidated.
	Invalidate(ctx context.Context, options InvalidateOptions) error
	
	// Clear removes all items from the cache.
	// Returns an error.
	Clear(ctx context.Context) error
}
```

## Memory (go-cache)

To create a new Memory (go-cache) store call `stash.NewMemory` and pass in a default expiry and default clean
up duration.

➡️ Click [here](https://github.com/lacuna-seo/stash/blob/dev/examples/memory.go) for an example.

```go
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

```

## Redis

To create a new Redis store call `stash.NewRedis` and pass in the redis options from `github.com/go-redis/redis/v8`
(ensure that it is imported) and a default expiry. 

➡️ Click [here](https://github.com/lacuna-seo/stash/blob/dev/examples/redis.go) for an example.

```go
provider := stash.NewRedis(redis.Options{
    Addr: "127.0.0.1:6379",
}, 5*time.Minute)

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
```

## Memcache

To create a new Memcache store call `stash.NewMemcache` and pass a slice of strings that correlate to a memcache 
server and a default expiry.

➡️ Click [here](https://github.com/lacuna-seo/stash/blob/dev/examples/memcache.go) for an example.

```go
provider := stash.NewMemcache([]string{"127.0.0.1:11211"}, 5*time.Minute)

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
```

## Tags

