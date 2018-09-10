package utils

import (
	"fmt"

	"github.com/satori/go.uuid"
)

//Create UUID
func CreateUUID() string {
	v, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return v.String()
}
