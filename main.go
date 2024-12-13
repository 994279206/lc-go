package main

import (
	"fmt"
	"lc/lru"
)

func main() {

	lruCache := lru.Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	fmt.Print(lruCache.Get(1))

}
