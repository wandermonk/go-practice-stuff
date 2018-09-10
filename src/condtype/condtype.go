package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	c.L.Lock()
	fmt.Println("Waiting......")
	for conditionTrue() == false {
		c.Wait()
	}
	fmt.Println("Done Waiting......")
	c.L.Unlock()

	go conditionTrue()
}

func conditionTrue() bool {
	fmt.Println("Inside conditionTrue() .........")
	var flag bool
	time.Sleep(1 * time.Millisecond)
	flag = true
	return flag
}
