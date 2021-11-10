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
* [Redis](https://github.com/go-redis/redis/v8) (go-redis/redis)
* [Memcache](https://github.com/bradfitz/gomemcache) (bradfitz/memcache)

## Install

`go get -u github.com/lacuna-seo/stash`

## Provider

A provider can is an interface that can be used to pass to `stash.Load`. It is used as a driver for
common methods between each memory store (Memory, Redis ot Memcache). 

It is the result of what is called by `NewMemory`, `NewRedis` or `NewMemcache`. Which can be pinged.
and validated. The methods are described below.


```go
type Provider interface {
    // Ping the store.
    Ping() error
    
    // Validate checks the environment for errors.
    Validate() error
    
    // Driver returns the store's name.
    Driver() string
    
    // Store returns the interface for use within
    // the cache.
    Store() store.StoreInterface
}
```

## Store

A store is what is used to interact with the cache driver. Items can retrieve, set, deleted, invalidated and
flushed. The methods are described below.

```go
type Store interface {
	// Get retrieves a specific item from the cache by key. Values are
	// automatically marshalled for use with Redis & Memcache.
	// Returns an error if the item could not be found
	// or unmarshalled.
	Get(ctx context.Context, key, v interface{}) error

	// Set stores a singular item in memory by key, value
	// and options (tags and expiration time). Values are automatically
	// marshalled for use with Redis & Memcache.
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

Cache invalidaton is hard. By using tags you are able to group cache items together.

```go
// Set a cache key with the value of 'stash' and a tag of 'category`.
// The cache item will be expired after one hour.
err := cache.Set(context.Background(), "key", []byte("stash"), stash.Options{
    Expiration: time.Hour * 1,
    Tags:       []string{"category"},
})
if err != nil {
    log.Fatalln(err)
}

// Invalidate all cache items that have a tag called 'category' 
// associated with them.
err = cache.Invalidate(context.Background(), stash.InvalidateOptions{
    Tags: []string{"category"},
})
if err != nil {
    log.Fatalln(err)
}
```

## Examples

To run the examples, clone the repo and run `make setup` and choose one of the following commands to run
an example with a particular store.

✅ **Memory**: `make memory:example`

✅ **Redis**: `make redis:example`

✅ **Memcache**: `make memcache:example`

## Credits

Thanks to [https://github.com/eko/gocache](https://github.com/eko/gocache)