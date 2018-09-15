package main

import (
	"fmt"
)

//Iterator allows to iterate a collection
type Iterator interface {
	Next() (element string, ok bool)
}

//Collection contains the index and list of strings
type Collection struct {
	index int
	list  []string
}

//Next gets the next element from the collection list
func (c *Collection) Next() (string, bool) {
	c.index++
	if c.index >= len(c.list) {
		return "INVALID_INDEX", false
	}
	return c.list[c.index], true
}

func newSlice(data []string) *Collection {
	return &Collection{
		index: -1,
		list:  data,
	}
}

func main() {
	var value string
	var ok bool
	var greetings Iterator
	greetings = newSlice([]string{"hi", "hello", "how are you"})
	value, ok = greetings.Next()
	for ok {
		fmt.Println(value)
		value, ok = greetings.Next()
	}
}
