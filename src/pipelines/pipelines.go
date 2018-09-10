package main

import (
	"fmt"
)

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 1, 1, 2, 4, 1}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 3))
	}
}
