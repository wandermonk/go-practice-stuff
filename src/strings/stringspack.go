package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := AllowedFile("photo.jp")
	fmt.Println(result1)
}

//AllowedFile checks for allowed files extensions for uploading into the system
func AllowedFile(filename string) bool {
	var ext string
	exts := []string{"png", "jpg", "jpeg"}
	if strings.ContainsAny(filename, ".") {
		ext = strings.ToLower(strings.Split(filename, ".")[1])
		return contains(exts, ext)
	}
	return false
}

func contains(exts []string, ext string) bool {
	for _, v := range exts {
		if ext == v {
			return true
		}
	}
	return false
}
