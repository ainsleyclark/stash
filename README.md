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

## Memory (go-cache)



## Redis

```go
provider := stash.NewRedis(redis.Options{
    Addr: "127.0.0.1:6379",
}, time.Hour*8)

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


```go
// Redis example for Stash.
func Redis() {
	// Create a provider, this could be Redis, Memcache
	// or In Memory (go-cache).
	provider := stash.NewRedis(redis.Options{
		Addr: "127.0.0.1:6379",
	}, time.Hour*8)

	// Create a new cache store by passing a provider.
	cache, err := stash.Load(provider)
	if err != nil {
		log.Fatalln(err)
	}

	// Set a cache item with key and value, with the expiration
	// time of one hour and a tag,
	err = cache.Set(context.Background(), "key", []byte("stash"), stash.Options{
		Expiration: time.Hour * 1,
		Tags:       []string{"tag"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Obtains the cache item by key which automatically unmarshalls
	// the value by passing a reference to the same type that
	// has been stored.
	var buf []byte
	err = cache.Get(context.Background(), "key", &buf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buf)) // Returns stash
}
```
