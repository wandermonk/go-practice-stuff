package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rock = iota
	paper
	scissors
)

const (
	player1 = "A"
	player2 = "B"
)

func main() {
	d := RockPaperScissor()
	if d == 0 {
		fmt.Println("TIE")
	} else if d%2 == 1 {
		fmt.Println("A WINS")
	} else if d%2 == 0 {
		fmt.Println("B WINS")
	}
}

//RockPaperScissor returns a result
func RockPaperScissor() int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(3)
	b := rand.Intn(3)
	d := (3 + a - b) % 3
	return d
}
