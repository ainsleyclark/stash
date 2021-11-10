// Copyright 2020 The Reddico Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/lacuna-seo/stash/examples"
	"log"
)

func main() {
	memory := flag.Bool("memory", false, "Pass to use the Memory example")
	redis := flag.Bool("redis", false, "Pass to use the Redis example")
	memcache := flag.Bool("memcache", false, "Pass to use the Memcache example")

	flag.Parse()

	if *memory {
		examples.Memory()
		return
	}

	if *redis {
		examples.Redis()
		return
	}

	if *memcache {
		examples.Memcache()
		return
	}

	log.Fatalln("No provider found use --memory, --redis or --memcache")
}
