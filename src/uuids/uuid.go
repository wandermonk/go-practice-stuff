package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	fmt.Println(CreateUUID())
}

func CreateUUID() string {
	v, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return v.String()
}
