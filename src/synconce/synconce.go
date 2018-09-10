package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once
var wg sync.WaitGroup

func main() {

	increment := func() {
		count++
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()

	fmt.Printf("The count is %d \n", count)

}
