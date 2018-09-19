package main

import (
	"fmt"
)

func main() {
	TableGenerator(1, 1)
	TableGenerator(2, 1)
}

//TableGenerator generates the table
func TableGenerator(number, limit int) {
	if limit == 11 {
		return
	}
	fmt.Printf("%d * %d = %d\n", number, limit, number*limit)
	TableGenerator(number, limit+1)
}
