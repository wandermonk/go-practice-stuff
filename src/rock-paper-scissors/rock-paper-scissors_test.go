package main

import (
	"testing"
)

func BenchmarkRPS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RockPaperScissor()
	}
}
