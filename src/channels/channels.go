package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello world"
	}()

	fmt.Println(<-ch)
}
