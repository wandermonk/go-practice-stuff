package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new instances")
			return struct{}{}
		},
	}

	pool.Get()
	instance := pool.Get()
	pool.Put(instance)
	pool.Get()

}
