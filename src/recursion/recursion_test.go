package recursion

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
	input := []int{1, 2, 3, 1, 3, 1, 4}
	for i := 0; i < b.N; i++ {
		Sum(input)
	}
}

func BenchmarkSum40(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	for i := 0; i < b.N; i++ {
		Sum(input)
	}
}
