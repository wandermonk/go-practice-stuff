package main

import (
	"fmt"
)

func main() {
	fmt.Println(FibMemoized(100000))
}

//Memoized is a type of function which takes int and return int
type Memoized func(int) int

//Memoize takes a memoized type and caches
func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

var fibMem Memoized

//FibMemoized memoized fib func
func FibMemoized(n int) int {
	fibMem = Memoize(func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fibMem(n-2) + fibMem(n-1)
	})
	return fibMem(n)
}
