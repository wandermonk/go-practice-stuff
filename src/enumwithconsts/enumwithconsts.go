package main

import (
	"fmt"
)

const (
	// UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	// TRACE logs everything
	TRACE // 1
	// INFO logs Info, Warnings and Errors
	INFO // 2
	// WARNING logs Warning and Errors
	WARNING // 3
	// ERROR just logs Errors
	ERROR // 4
)

// Level holds the log level.
type Level int

func SetLogLevel(level Level) {
	switch level {
	case TRACE:
		fmt.Println("trace")
		return

	case INFO:
		fmt.Println("info")
		return

	case WARNING:
		fmt.Println("warning")
		return
	case ERROR:
		fmt.Println("error")
		return

	default:
		fmt.Println("default")
		return

	}
}

func main() {

	SetLogLevel(INFO)

}
