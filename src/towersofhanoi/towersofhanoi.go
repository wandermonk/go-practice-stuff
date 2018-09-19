package main

import (
	"fmt"
)

//TowersOfHanoi is a function which moves the disks using pegs
func TowersOfHanoi(source, destination, extra string, disks int) {
	if disks <= 0 {
		return
	}
	TowersOfHanoi(source, extra, destination, disks-1)
	fmt.Printf("Move Disks %d from %s to %s\n", disks, source, destination)
	TowersOfHanoi(extra, destination, source, disks-1)
}

func main() {
	TowersOfHanoi("s", "d", "e", 3)
}
