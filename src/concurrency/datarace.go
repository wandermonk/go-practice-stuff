package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Counter struct {
	m     sync.Mutex
	count int
}

func (c *Counter) increment() {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	fmt.Printf("The count is %d\n", c.count)
}

func main() {
	wg.Add(2)
	counter := &Counter{}
	go counter.increment()
	go counter.increment()
	wg.Wait()
}
