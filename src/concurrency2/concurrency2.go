package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// var input chan bool

func main() {
	wg.Add(1)
	// input = make(chan bool)
	go func() {
		defer wg.Done()
		var count int
		var flag bool
		for i := 0; i < 1000; i++ {
			count++
		}
		if count == 1000 {
			fmt.Printf("The count is %d\n", count)
			flag = true
		}
		fmt.Println(flag)
		// input <- flag
		// close(input)
	}()

	// val, ok := <-input
	// if ok {
	// 	fmt.Println(val)
	// }

	wg.Wait()
}
