package main

import (
	"fmt"
)

type Enum interface {
	name() string
	ordinal() int
	value() *[]string
}

var gender = []string{"male", "female"}

type Gender uint

func (g Gender) name() string {
	return gender[g]
}

func (g Gender) ordinal() int {
	return int(g)
}

func (g Gender) value() *[]string {
	return &gender
}

const (
	Male Gender = iota
	Female
)

func main() {
	fmt.Println("male")
}
