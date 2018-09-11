package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	d := RockPaperScissor()
	fmt.Println(d)
}

//RockPaperScissor returns a result
func RockPaperScissor() string {
	var result string
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(3)
	b := rand.Intn(3)
	d := (3 + a - b) % 3
	if d == 0 {
		result = "TIE"
	} else if d%2 == 1 {
		result = "A WINS"
	} else if d%2 == 0 {
		result = "B WINS"
	}
	return result
}
